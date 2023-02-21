package hap

import (
	"github.com/brutella/hap/log"
	"github.com/brutella/hap/tlv8"

	"net/http"
	"reflect"
)

type pairingPayload struct {
	Identifier string `tlv8:"1"`
	PublicKey  []byte `tlv8:"3"`
	Permission byte   `tlv8:"11"`
}

func (srv *Server) pairings(res http.ResponseWriter, req *http.Request) {
	if !srv.IsAuthorized(req) {
		log.Info.Printf("request from %s not authorized\n", req.RemoteAddr)
		JsonError(res, JsonStatusInsufficientPrivileges)
		return
	}

	ss, err := srv.getSession(req.RemoteAddr)
	if err != nil {
		log.Info.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		tlv8Error(res, M2, TlvErrorUnknown)
		return
	}

	d := struct {
		Method     byte   `tlv8:"0"`
		Identifier string `tlv8:"1"`
		PublicKey  []byte `tlv8:"3,optional"`
		Permission byte   `tlv8:"11,optional"`
		State      byte   `tlv8:"6"`
	}{}

	if err := tlv8.UnmarshalReader(req.Body, &d); err != nil {
		log.Info.Println("tlv8:", err)
		res.WriteHeader(http.StatusBadRequest)
		tlv8Error(res, M2, TlvErrorUnknown)
		return
	}

	switch d.Method {
	case MethodAddPairing:
		log.Debug.Println("add pairing", d.Identifier)

		if ss.Pairing.Permission != PermissionAdmin {
			log.Info.Println("operation not allowed for non-admin controllers")
			tlv8Error(res, M2, TlvErrorAuthentication)
			return
		}

		p, err := srv.st.Pairing(d.Identifier)
		if err != nil {
			p = Pairing{
				Name:       d.Identifier,
				PublicKey:  d.PublicKey,
				Permission: d.Permission,
			}
		} else {
			if !reflect.DeepEqual(p.PublicKey, d.PublicKey) {
				log.Info.Println("invalid public key")
				tlv8Error(res, M2, TlvErrorUnknown)
				return
			}
			// Update permission
			p.Permission = d.Permission
		}

		err = srv.savePairing(p)
		if err != nil {
			log.Info.Println(err)
			tlv8Error(res, M2, TlvErrorUnknown)
			return
		}

		resp := struct {
			State byte `tlv8:"6"`
		}{
			State: M2,
		}
		tlv8OK(res, resp)

	case MethodDeletePairing:
		log.Debug.Println("delete pairing", d.Identifier)

		if ss.Pairing.Permission != PermissionAdmin {
			log.Info.Println("operation not allowed for non-admin controllers")
			tlv8Error(res, M2, TlvErrorAuthentication)
			return
		}

		p, err := srv.st.Pairing(d.Identifier)
		if err != nil {
			log.Info.Println(err)
			tlv8Error(res, M2, TlvErrorUnknown)
			return
		}

		if err = srv.deletePairing(p); err != nil {
			log.Info.Println(err)
			tlv8Error(res, M2, TlvErrorUnknown)
			return
		}

		resp := struct {
			State byte `tlv8:"6"`
		}{
			State: 2,
		}
		tlv8OK(res, resp)

		// If no admin controller is paired anymore,
		// close all connections and delete all pairings
		if !srv.pairedWithAdmin() {
			for addr, conn := range conns() {
				log.Debug.Println("Closing connection to", addr)
				conn.Close()
			}
			srv.deleteAllPairings()
		}

		// Close connection of deleted controller
		for addr, conn := range conns() {
			ss, err := srv.getSession(addr)
			if err != nil {
				log.Debug.Println("no session for", addr, err)
				continue
			}
			if ss.Pairing.Name == p.Name {
				log.Debug.Println("closing connection of removed controller", d.Identifier)
				conn.Close()
			}
		}

	case MethodListPairings:
		log.Debug.Println("list pairings")
		ps := srv.st.Pairings()
		resp := make([]pairingPayload, len(ps))
		for i, p := range ps {
			resp[i] = pairingPayload{
				Identifier: p.Name,
				PublicKey:  p.PublicKey,
				Permission: p.Permission,
			}
		}
		tlv8OK(res, resp)
	}
}
