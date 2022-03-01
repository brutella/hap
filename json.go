package hap

import (
	"encoding/json"
	"net/http"
)

// JsonOK sends an HTTP 200 (ok) response.
func JsonOK(res http.ResponseWriter, body interface{}) error {
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}

	res.WriteHeader(http.StatusOK)
	wr := NewChunkedWriter(res, 2048)
	_, err = wr.Write(b)
	return err
}

// JsonMultiStatus sends an HTTP 207 (multi status) response.
func JsonMultiStatus(res http.ResponseWriter, body interface{}) error {
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}

	res.WriteHeader(http.StatusMultiStatus)
	wr := NewChunkedWriter(res, 2048)
	_, err = wr.Write(b)
	return err
}

// JsonErrors sends an HTTP 500 (bad request) response including the status in the body.
func JsonError(res http.ResponseWriter, status int) error {
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
