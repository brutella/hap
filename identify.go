package hap

import (
	"github.com/brutella/hap/log"

	"net/http"
)

func (srv *Server) identify(res http.ResponseWriter, req *http.Request) {
	if !srv.IsAuthorized(req) {
		log.Info.Printf("request from %s not authorized\n", req.RemoteAddr)
		JsonError(res, JsonStatusInsufficientPrivileges)
		return
	}

	if srv.a.IdentifyFunc != nil {
		srv.a.IdentifyFunc(req)
	}

	JsonOK(res, struct{}{})
}
