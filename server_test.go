package hap

import (
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"

	"net/http"
	"net/http/httptest"
	"testing"
)

// TestConfigHash tests if the server updates the config hash
// when the accessory configuration changed.
func TestConfigHash(t *testing.T) {
	a := accessory.New(accessory.Info{Name: "ABC"}, accessory.TypeOutlet)
	a.AddS(service.NewOutlet().S)

	st := NewMemStore()
	s, err := NewServer(st, a)
	if err != nil {
		t.Fatal(err)
	}

	v1 := s.version

	// Change the structure of the accessory by adding another service.
	// The server has to update the version to 2.
	a.AddS(service.NewSwitch().S)
	s, err = NewServer(st, a)
	if err != nil {
		t.Fatal(err)
	}

	if is, want := s.version, v1+1; is != want {
		t.Fatalf("%v != %v", is, want)
	}
}

func TestIdentify(t *testing.T) {
	a := accessory.New(accessory.Info{Name: "ABC"}, accessory.TypeOutlet)

	s, err := NewServer(NewMemStore(), a)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodPost, "/identify", nil)
	w := httptest.NewRecorder()

	var identified bool
	a.IdentifyFunc = func(r *http.Request) {
		if is, want := r, req; is != want {
			t.Fatalf("%v != %v", is, want)
		}
		identified = true
	}

	s.Identify(w, req)

	r := w.Result()
	if is, want := r.StatusCode, http.StatusOK; is != want {
		t.Fatalf("%v != %v", is, want)
	}

	if is, want := identified, true; is != want {
		t.Fatalf("%v != %v", is, want)
	}
}
