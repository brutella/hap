package hap

import (
	"sync"
	"time"

	"github.com/brutella/dnssd"
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/characteristic"
	"github.com/brutella/hap/log"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/xiam/to"
	godiacritics "gopkg.in/Regis24GmbH/go-diacritics.v2"

	"bytes"
	"context"
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"
	"net"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

// A server handles incoming HTTP request for an accessory.
// The server uses dnssd to announce the accessory on the local network.
type Server struct {
	// Pin specifies the pincode used to pair
	// with the accessory.
	Pin string

	// Addr specifies the tcp address for the server
	// to listen to in form of "host:port".
	// If empty, a random port is used.
	Addr string

	// Ifaces specifies at which interface the
	// associated dnssd service is announced.
	Ifaces []string

	MfiCompliant bool   // default false
	Protocol     string // default "1.0"
	SetupId      string
	Key          KeyPair // public and private key (generated and stored on disk)

	st *storer        // stores data
	ss *http.Server   // http server
	a  *accessory.A   // main accessory
	as []*accessory.A // bridged accessories

	version uint16 // version of accessory content – relates to configHash
	uuid    string // internal identifier (generated and stored on disk)

	port int // listen port (can be different than in Addr)
	ln   *net.TCPListener

	// for dnssd stuff
	responder dnssd.Responder
	handle    dnssd.ServiceHandle

	mux  *sync.Mutex
	sess map[string]interface{}
	cons map[string]*conn
}

// A ServeMux lets you attach handlers to http url paths.
type ServeMux interface {
	// Handle registers the handler for the given pattern.
	Handle(pattern string, handler http.Handler)
	// HandleFuncs registers the handler function for the given pattern.
	HandleFunc(pattern string, handler http.HandlerFunc)
	// Mount attaches another http.Handler along ./pattern/*
	Mount(pattern string, handler http.Handler)
}

// NewServer returns a new server given a store (to persist data) and accessories.
// If more than one accessory is added to the server, *a* acts as a bridge.
func NewServer(store Store, a *accessory.A, as ...*accessory.A) (*Server, error) {
	r := chi.NewRouter()
	r.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: log.Debug, NoColor: true}))

	st := &storer{store}
	if err := migrate(st); err != nil {
		log.Info.Panic(err)
	}

	s := &Server{
		st:   st,
		a:    a,
		as:   as,
		mux:  &sync.Mutex{},
		sess: make(map[string]interface{}),
		cons: make(map[string]*conn),
	}
	s.ss = &http.Server{
		Handler:   r,
		ConnState: s.connStateEvent,
	}

	// Load the stored uuid or generate a new one.
	if s.uuid == "" {
		uuid, err := s.st.Get("uuid")
		if err != nil {
			uuid = []byte(mac48Address(randHex()))
			if err := s.st.Set("uuid", uuid); err != nil {
				return nil, err
			}
		}

		s.uuid = string(uuid)
	}

	// Load the stored version or set to 1.
	if s.version == 0 {
		b, err := s.st.Get("version")
		if err == nil {
			s.version = uint16(to.Uint64(string(b)))
		} else {
			s.version = 1
		}
	}

	arr := []*accessory.A{a}
	arr = append(arr, as[:]...)
	if err := s.add(arr); err != nil {
		return nil, err
	}

	// Group handlers for tlv8 and json encoded content.
	r.Group(func(r chi.Router) {
		r.Use(middleware.SetHeader("Content-Type", HTTPContentTypePairingTLV8))
		r.Post("/pair-setup", s.pairSetup)
		r.Post("/pair-verify", s.pairVerify)
		r.Post("/identify", s.identify)
		r.Post("/pairings", s.pairings)
	})

	// The json encoded content is encrypted. The encryption keys
	// are stored in a session. The de-/encryption is done by a Conn.
	r.Group(func(r chi.Router) {
		r.Use(middleware.SetHeader("Content-Type", HTTPContentTypeHAPJson))
		r.Get("/accessories", s.getAccessories)
		r.Get("/characteristics", s.getCharacteristics)
		r.Put("/characteristics", s.putCharacteristics)
		r.Put("/prepare", s.prepareCharacteristics)
	})

	return s, nil
}

// ServeMux returns the http handler.
func (s *Server) ServeMux() ServeMux {
	return s.ss.Handler.(*chi.Mux)
}

// IsAuthorized returns true if the provided
// request is authorized to access accessory data.
func (s *Server) IsAuthorized(request *http.Request) bool {
	ss, _ := s.getSession(request.RemoteAddr)
	return ss != nil
}

func (s *Server) TimedWrite(request *http.Request) *TimedWrite {
	if ss, _ := s.getSession(request.RemoteAddr); ss != nil {
		return ss.twr
	}

	return nil
}

func (s *Server) SetTimedWrite(ttl, pid uint64, request *http.Request) {
	if ss, _ := s.getSession(request.RemoteAddr); ss != nil {
		t := time.Now().Add(time.Duration(ttl) * time.Millisecond)
		ss.twr = &TimedWrite{t, pid}
	}
}

func (s *Server) DelTimedWrite(request *http.Request) {
	if ss, _ := s.getSession(request.RemoteAddr); ss != nil {
		ss.twr = nil
	}
}

// IsPaired returns true if the server is paired with a client (iOS).
func (s *Server) IsPaired() bool {
	return len(s.st.Pairings()) > 0
}

// ListenAndServe starts the server.
func (s *Server) ListenAndServe(ctx context.Context) error {
	err := s.prepare()
	if err != nil {
		return err
	}

	return s.listenAndServe(ctx)
}

func (s *Server) listenAndServe(ctx context.Context) error {
	// Listen with a tcp socket on a given addr/port.
	tcpLn, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	ln := &listener{tcpLn.(*net.TCPListener)}

	// Get the port from the listener address because it
	// it might be different than specified in Port.
	_, port, _ := net.SplitHostPort(ln.Addr().String())
	i, err := strconv.Atoi(port)
	if err != nil {
		return err
	}
	s.port = i

	// Announce the server using dnssd.
	resp, err := dnssd.NewResponder()
	if err != nil {
		return fmt.Errorf("dnssd: %s", err)
	}
	s.responder = resp

	service, err := s.service()
	if err != nil {
		return fmt.Errorf("dnssd: %s", err)
	}

	h, err := resp.Add(service)
	if err != nil {
		return err
	}
	s.handle = h

	dnsCtx, dnsCancel := context.WithCancel(ctx)
	defer dnsCancel()

	dnsStop := make(chan struct{})
	go func() {
		resp.Respond(dnsCtx)
		log.Debug.Println("dnssd responder stopped")
		dnsStop <- struct{}{}
	}()

	log.Debug.Println("listening at", ln.Addr())

	serverCtx, serverCancel := context.WithCancel(ctx)
	defer serverCancel()

	serverStop := make(chan struct{})
	go func() {
		<-serverCtx.Done()
		s.ss.Close()
		ln.Close()
		log.Debug.Println("http server stopped")
		serverStop <- struct{}{}
	}()

	err = s.ss.Serve(ln)
	<-dnsStop
	<-serverStop

	return err
}

func (s *Server) add(as []*accessory.A) error {
	aid := uint64(1)
	for _, a := range as {
		if a.Name() == "" {
			return errors.New("invalid accessory name")
		}

		if a.Id == 0 {
			a.Id = aid
			aid++
		}

		var iid uint64 = 1
		for _, s := range a.Ss {
			s.Id = iid
			iid++

			for _, c := range s.Cs {
				// Create a local variable before
				// capturing them in a function.
				a := a

				c.Id = iid
				iid++

				// If the value of a characteristic changes, we notify all connected clients.
				// The identify characteristic is a special case where we all accessory.IdentifyFunc.
				if c.Type == characteristic.TypeIdentify {
					c.OnCValueUpdate(func(c *characteristic.C, new, old interface{}, req *http.Request) {
						if b, ok := new.(bool); ok && b && a.IdentifyFunc != nil {
							a.IdentifyFunc(req)
						}
					})
				} else {
					c.OnCValueUpdate(func(c *characteristic.C, new, old interface{}, req *http.Request) {
						// send notification to all subscribed clients
						sendNotification(a, c, req)
					})
				}
			}
		}
	}

	// The server keeps track of previously published accessories.
	// If the accessory changed (added service or characteristics)
	// from last time, we have to update the version flag.
	var oldHash, newHash []byte

	if b, err := s.st.Get("configHash"); err == nil && len(b) > 0 {
		oldHash = b
	}
	newHash = configHash(as)
	if !reflect.DeepEqual(oldHash, newHash) {
		s.version += 1
		s.st.Set("version", []byte(fmt.Sprintf("%d", s.version)))
		s.st.Set("configHash", newHash)
	}

	return nil
}

func (s *Server) prepare() error {
	if allZero(s.Key.Public[:]) || allZero(s.Key.Private[:]) {
		// Load keypair or generate a new one.
		keypair, err := s.st.KeyPair()
		if err != nil {
			keypair, err := generateKeyPair()
			if err != nil {
				return fmt.Errorf("generating keypair failed: %v", err)
			}
			s.Key = keypair
			if err := s.st.SaveKeyPair(keypair); err != nil {
				return fmt.Errorf("saving keypair failed: %v", err)
			}
		} else {
			s.Key = keypair
		}
	}

	if s.Pin == "" {
		s.Pin = "00102003" // default pincode
	}

	if s.Protocol == "" {
		s.Protocol = "1.0"
	}

	if len(s.Pin) != 8 {
		return fmt.Errorf("invald pin length %d", len(s.Pin))
	} else if _, found := InvalidPins[s.Pin]; found {
		return fmt.Errorf("insecure pin %s", s.Pin)
	}

	return nil
}

func (s *Server) connStateEvent(conn net.Conn, event http.ConnState) {
	if event == http.StateClosed {
		addr := conn.RemoteAddr().String()
		s.mux.Lock()
		delete(s.sess, addr)
		delete(s.cons, addr)
		s.mux.Unlock()
	}
}

func (s *Server) getSession(addr string) (*session, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	if v, ok := s.sess[addr]; ok {
		if s, ok := v.(*session); ok {
			return s, nil
		}
		return nil, fmt.Errorf("unexpected session %T", v)
	}

	return nil, fmt.Errorf("no session for %s", addr)
}

func (s *Server) getPairVerifySession(addr string) (*pairVerifySession, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	if v, ok := s.sess[addr]; ok {
		if s, ok := v.(*pairVerifySession); ok {
			return s, nil
		}
		return nil, fmt.Errorf("unexpected session %T", v)
	}

	return nil, fmt.Errorf("no session for %s", addr)
}

func (s *Server) getPairSetupSession(addr string) (*pairSetupSession, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	if v, ok := s.sess[addr]; ok {
		if s, ok := v.(*pairSetupSession); ok {
			return s, nil
		}
		return nil, fmt.Errorf("unexpected session %T", v)
	}

	return nil, fmt.Errorf("no session for %s", addr)
}

func (s *Server) setSession(addr string, v interface{}) {
	s.mux.Lock()
	s.sess[addr] = v
	s.mux.Unlock()
}

func (s *Server) sessions() map[string]interface{} {
	copy := map[string]interface{}{}
	s.mux.Lock()
	for k, v := range s.sess {
		copy[k] = v
	}
	s.mux.Unlock()

	return copy
}

func (s *Server) savePairing(p Pairing) error {
	err := s.st.SavePairing(p)
	if err != nil {
		return err
	}

	s.updateTxtRecords()
	return nil
}

func (s *Server) deletePairing(p Pairing) error {
	err := s.st.DeletePairing(p.Name)
	if err != nil {
		return err
	}

	s.updateTxtRecords()
	return nil
}

func (s *Server) deleteAllPairings() {
	for _, p := range s.st.Pairings() {
		s.st.DeletePairing(p.Name)
	}
	s.updateTxtRecords()
}

func (s *Server) pairedWithAdmin() bool {
	for _, p := range s.st.Pairings() {
		if p.Permission == PermissionAdmin {
			return true
		}
	}

	return false
}

func (s *Server) txtRecords() map[string]string {
	return map[string]string{
		"pv": s.Protocol,
		"id": s.uuid,
		"c#": fmt.Sprintf("%d", s.version),
		"s#": "1",
		"sf": fmt.Sprintf("%d", to.Int64(!s.IsPaired())),
		"ff": fmt.Sprintf("%d", to.Int64(s.MfiCompliant)),
		"md": s.a.Name(),
		"ci": fmt.Sprintf("%d", s.a.Type),
		"sh": s.setupHash(),
	}
}

func (s *Server) setupHash() string {
	hashvalue := fmt.Sprintf("%s%s", s.SetupId, s.uuid)
	sum := sha512.Sum512([]byte(hashvalue))
	// use only first 4 bytes
	code := []byte{sum[0], sum[1], sum[2], sum[3]}
	encoded := base64.StdEncoding.EncodeToString(code)
	return encoded
}

func (s *Server) updateTxtRecords() {
	if s.handle != nil {
		s.handle.UpdateText(s.txtRecords(), s.responder)
	}
}

func (s *Server) service() (dnssd.Service, error) {
	// 2016-03-14(brutella): Replace whitespaces (" ") from service name
	// with underscores ("_")to fix invalid http host header field value
	// produces by iOS.
	//
	// [Radar] http://openradar.appspot.com/radar?id=4931940373233664
	stripped := strings.Replace(s.a.Info.Name.Value(), " ", "_", -1)
	cfg := dnssd.Config{
		Name:   normalize(stripped),
		Type:   "_hap._tcp",
		Domain: "local",
		Host:   strings.Replace(s.uuid, ":", "", -1), // use the id (without the colons) to get unique hostnames
		Text:   s.txtRecords(),
		Port:   s.port,
		Ifaces: s.Ifaces,
	}

	return dnssd.NewService(cfg)
}

var InvalidPins = map[string]bool{
	"00000000": true,
	"11111111": true,
	"22222222": true,
	"33333333": true,
	"44444444": true,
	"55555555": true,
	"66666666": true,
	"77777777": true,
	"88888888": true,
	"99999999": true,
	"12345678": true,
	"87654321": true,
}

func (s *Server) fmtPin() string {
	runes := bytes.Runes([]byte(s.Pin))
	first := string(runes[:3])
	second := string(runes[3:5])
	third := string(runes[5:])
	return first + "-" + second + "-" + third
}

func normalize(str string) string {
	return godiacritics.Normalize(str)
}

func allZero(s []byte) bool {
	for _, v := range s {
		if v != 0 {
			return false
		}
	}
	return true
}
