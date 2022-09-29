package hap

import (
	"github.com/brutella/hap/chacha20poly1305"
	"github.com/brutella/hap/curve25519"
	"github.com/brutella/hap/ed25519"
	"github.com/brutella/hap/hkdf"
	"github.com/brutella/hap/log"
	"github.com/brutella/hap/tlv8"

	"net/http"
)

type pairVerifyPayload struct {
	Method        byte   `tlv8:"0,optional"`
	Identifier    string `tlv8:"1,optional"`
	PublicKey     []byte `tlv8:"3,optional"`
	EncryptedData []byte `tlv8:"5,optional"`
	State         byte   `tlv8:"6,optional"`
	Signature     []byte `tlv8:"10,optional"`
}

type pairVerifySession struct {
	OtherPublicKey []byte
	PublicKey      [32]byte
	PrivateKey     [32]byte
	SharedKey      [32]byte
	EncryptionKey  [32]byte
}

func (srv *Server) pairVerify(res http.ResponseWriter, req *http.Request) {
	data := pairVerifyPayload{}
	if err := tlv8.UnmarshalReader(req.Body, &data); err != nil {
		log.Info.Println("tlv8:", err)
		tlv8Error(res, data.State+1, TlvErrorUnknown)
		return
	}

	switch data.Method {
	case MethodPair:
		switch data.State {
		case M1:
			srv.pairVerifyM1(res, req, data)
		case M3:
			srv.pairVerifyM3(res, req, data)
		default:
			log.Info.Println("invalid state", data.State)
			res.WriteHeader(http.StatusBadRequest)
			tlv8Error(res, data.State+1, TlvErrorUnknown)
		}
	default:
		log.Info.Println("pair verify: invalid method", data.Method)
		res.WriteHeader(http.StatusBadRequest)
		tlv8Error(res, 0, TlvErrorInvalidRequest)
	}

}

func (srv *Server) pairVerifyM1(res http.ResponseWriter, req *http.Request, data pairVerifyPayload) {
	var otherPublicKey [32]byte
	copy(otherPublicKey[:], data.PublicKey)

	// Generate the key pair.
	publicKey, privateKey := curve25519.GenerateKeyPair()
	sharedKey := curve25519.SharedSecret(privateKey, otherPublicKey)
	encKey, err := hkdf.Sha512(sharedKey[:], []byte("Pair-Verify-Encrypt-Salt"), []byte("Pair-Verify-Encrypt-Info"))
	if err != nil {
		log.Info.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		tlv8Error(res, M2, TlvErrorUnknown)
		return
	}

	var buf []byte
	buf = append(buf, publicKey[:]...)
	buf = append(buf, srv.uuid...)
	buf = append(buf, data.PublicKey[:]...)
	signature, err := ed25519.Signature(srv.Key.Private[:], buf)
	if err != nil {
		log.Info.Println(err)
		tlv8Error(res, M2, TlvErrorUnknown)
		return
	}

	enData := struct {
		Identifier string `tlv8:"1"`
		Signature  []byte `tlv8:"10"`
	}{
		Identifier: srv.uuid,
		Signature:  signature,
	}

	b, err := tlv8.Marshal(enData)
	if err != nil {
		log.Info.Println("tlv8:", err)
		res.WriteHeader(http.StatusBadRequest)
		tlv8Error(res, M2, TlvErrorUnknown)
		return
	}

	encBuf, mac, _ := chacha20poly1305.EncryptAndSeal(encKey[:], []byte("PV-Msg02"), b, nil)
	resp := struct {
		State         byte   `tlv8:"6"`
		PublicKey     []byte `tlv8:"3"`
		EncryptedData []byte `tlv8:"5"`
	}{
		State:         M2,
		PublicKey:     publicKey[:],
		EncryptedData: append(encBuf, mac[:]...),
	}
	tlv8OK(res, resp)

	// Save the keys in a session and store the session for the request.
	ses := &pairVerifySession{
		OtherPublicKey: data.PublicKey,
		PublicKey:      publicKey,
		PrivateKey:     privateKey,
		SharedKey:      sharedKey,
		EncryptionKey:  encKey,
	}
	srv.setSession(req.RemoteAddr, ses)
}

func (srv *Server) pairVerifyM3(res http.ResponseWriter, req *http.Request, data pairVerifyPayload) {
	// Get the session for the request.
	ses, err := srv.getPairVerifySession(req.RemoteAddr)
	if err != nil {
		log.Info.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		tlv8Error(res, M4, TlvErrorUnknown)
		return
	}

	msg := data.EncryptedData[:len(data.EncryptedData)-16]
	var mac [16]byte
	copy(mac[:], data.EncryptedData[len(msg):]) // 16 byte (MAC)

	enc, err := chacha20poly1305.DecryptAndVerify(ses.EncryptionKey[:], []byte("PV-Msg03"), msg, mac, nil)
	if err != nil {
		log.Info.Println(err)
		tlv8Error(res, M4, TlvErrorAuthentication)
		return
	}

	encData := pairVerifyPayload{}
	if err := tlv8.Unmarshal(enc, &encData); err != nil {
		log.Info.Println("tlv8:", err)
		tlv8Error(res, M4, TlvErrorUnknown)
		return
	}

	pairing, err := srv.st.Pairing(encData.Identifier)
	if err != nil {
		log.Info.Printf("not paired with %s yet\n", encData.Identifier)
		tlv8Error(res, M4, TlvErrorAuthentication)
		return
	}

	var buf []byte
	buf = append(buf, ses.OtherPublicKey[:]...)
	buf = append(buf, []byte(encData.Identifier)...)
	buf = append(buf, ses.PublicKey[:]...)

	if !ed25519.ValidateSignature(pairing.PublicKey[:], buf, encData.Signature) {
		log.Info.Println("signature is invalid")
		tlv8Error(res, M4, TlvErrorUnknownPeer)
		return
	}

	resp := struct {
		State byte `tlv8:"6"`
	}{
		State: M4,
	}
	tlv8OK(res, resp)

	// Store the negotiated keys in a session.
	ss, err := newSession(ses.SharedKey, pairing)
	if err != nil {
		log.Info.Println(err)
		return
	}

	// Store the session for the request.
	srv.setSession(req.RemoteAddr, ss)

	conn := getConn(req)
	if conn == nil {
		log.Info.Printf("no connection for %s\n", req.RemoteAddr)
		return
	}

	// Upgrade the connection to use encryption.
	conn.Upgrade(ss)
}
