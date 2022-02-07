package hap

import (
	"github.com/brutella/hap/accessory"

	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAccessories(t *testing.T) {
	a := accessory.NewOutlet(accessory.Info{
		Name:         "Outlet",
		SerialNumber: "1234",
		Model:        "a",
		Manufacturer: "Matthias",
		Firmware:     "1.0",
	})

	a.Outlet.On.Val = true

	s, err := NewServer(NewMemStore(), a.A)
	if err != nil {
		t.Fatal(err)
	}

	// fake pairing
	p := Pairing{
		Name: "unit test",
	}
	if err := s.savePairing(p); err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodGet, "/accessories", new(bytes.Buffer))
	w := httptest.NewRecorder()

	s.GetAccessories(w, req)

	r := w.Result()
	if is, want := r.StatusCode, http.StatusOK; is != want {
		t.Fatalf("%v != %v", is, want)
	}

	b, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}

	if is, want := string(b), ""; is != want {
		t.Fatalf("%v != %v", is, want)
	}
}
