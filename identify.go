package hap

import (
	"net/http"
)

func (srv *Server) identify(res http.ResponseWriter, req *http.Request) {
	if srv.isPaired() {
		res.WriteHeader(http.StatusBadRequest)
		jsonError(res, JsonStatusInsufficientPrivileges)
		return
	}

	if srv.a.IdentifyFunc != nil {
		srv.a.IdentifyFunc(req)
	}

	jsonOK(res, struct{}{})
}
