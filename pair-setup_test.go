package hap

import (
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/tlv8"

	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPairSetup(t *testing.T) {
	a := accessory.New(accessory.Info{Name: "Outlet"}, accessory.TypeOutlet)
	s, err := NewServer(NewMemStore(), a)
	s.Pin = "00102003"

	if err != nil {
		t.Fatal(err)
	}

	d := PairSetupStep1Payload{
		State: Step1,
	}

	b, _ := tlv8.Marshal(d)
	req := httptest.NewRequest(http.MethodGet, "/pair-setup", bytes.NewBuffer(b))
	w := httptest.NewRecorder()

	s.PairSetup(w, req)

	r := w.Result()

	if is, want := r.StatusCode, http.StatusOK; is != want {
		t.Fatalf("%v != %v", is, want)
	}

	resp := PairSetupStep2Payload{}
	tlv8.UnmarshalReader(w.Body, &resp)

	if is, want := resp.State, Step2; is != want {
		t.Fatalf("%v != %v", is, want)
	}

	if is, want := len(resp.Salt), 16; is != want {
		t.Fatalf("%v != %v", is, want)
	}

	if is, want := len(resp.PublicKey), 384; is != want {
		t.Fatalf("%v != %v", is, want)
	}
}
