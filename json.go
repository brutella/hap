package hap

import (
	"encoding/json"
	"net/http"
)

func jsonOK(res http.ResponseWriter, body interface{}) error {
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}

	res.WriteHeader(http.StatusOK)
	wr := newChunkedWriter(res, 2048)
	_, err = wr.Write(b)
	return err
}

func jsonMultiStatus(res http.ResponseWriter, body interface{}) error {
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}

	res.WriteHeader(http.StatusMultiStatus)
	wr := newChunkedWriter(res, 2048)
	_, err = wr.Write(b)
	return err
}

func jsonError(res http.ResponseWriter, status int) error {
	resp := struct {
		Status int `json:"status"`
	}{
		Status: status,
	}

	b, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	res.WriteHeader(http.StatusBadRequest)
	_, err = res.Write(b)
	return err
}

func toJSON(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}

	return string(b)
}
