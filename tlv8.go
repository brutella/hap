package hap

import (
	"github.com/brutella/hap/tlv8"

	"net/http"
)

func tlv8OK(res http.ResponseWriter, body interface{}) error {
	b, err := tlv8.Marshal(body)
	if err != nil {
		return err
	}

	_, err = res.Write(b)
	return err
}

func tlv8Error(res http.ResponseWriter, state byte, status byte) error {
	resp := struct {
		State  byte `tlv8:"6"`
		Status byte `tlv8:"7"`
	}{
		State:  state,
		Status: status,
	}

	b, err := tlv8.Marshal(resp)
	if err != nil {
		return err
	}

	_, err = res.Write(b)
	return err
}
