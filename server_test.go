package hap

import (
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"

	"bytes"
	"fmt"
	"io/ioutil"
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

	s.setSession(req.RemoteAddr, &session{})

	var identified bool
	a.IdentifyFunc = func(r *http.Request) {
		if is, want := r, req; is != want {
			t.Fatalf("%v != %v", is, want)
		}
		identified = true
	}

	s.identify(w, req)

	r := w.Result()
	if is, want := r.StatusCode, http.StatusNoContent; is != want {
		t.Fatalf("%v != %v", is, want)
	}

	if is, want := identified, true; is != want {
		t.Fatalf("%v != %v", is, want)
	}
}

func TestSetValueRequestSuccess(t *testing.T) {
	a := accessory.NewOutlet(accessory.Info{Name: "ABC"})

	s, err := NewServer(NewMemStore(), a.A)
	if err != nil {
		t.Fatal(err)
	}

	body := fmt.Sprintf("{\"characteristics\":[{\"aid\":%d,\"iid\":%d,\"value\":true}]}", a.Id, a.Outlet.On.Id)
	req := httptest.NewRequest(http.MethodPut, "/characteristics", bytes.NewBuffer([]byte(body)))
	w := httptest.NewRecorder()

	s.setSession(req.RemoteAddr, &session{})

	var setValueRequestFunc, onValueUpdateFunc bool
	a.Outlet.On.SetValueRequestFunc = func(v interface{}, r *http.Request) (interface{}, int) {
		if is, want := v.(bool), true; is != want {
			t.Fatalf("%v != %v", is, want)
		}
		setValueRequestFunc = true

		return v, 0
	}

	a.Outlet.On.OnValueUpdate(func(new bool, old bool, r *http.Request) {
		if is, want := new, true; is != want {
			t.Fatalf("%v != %v", is, want)
		}

		if is, want := old, false; is != want {
			t.Fatalf("%v != %v", is, want)
		}

		onValueUpdateFunc = true
	})

	s.ss.Handler.ServeHTTP(w, req)

	r := w.Result()
	if is, want := r.StatusCode, http.StatusNoContent; is != want {
		t.Fatalf("%v != %v", is, want)
	}

	if is, want := setValueRequestFunc, true; is != want {
		t.Fatalf("%v != %v", is, want)
	}

	if is, want := onValueUpdateFunc, true; is != want {
		t.Fatalf("%v != %v", is, want)
	}

	if is, want := a.Outlet.On.Value(), true; is != want {
		t.Fatalf("%v != %v", is, want)
	}
}

func TestSetValueRequestFailure(t *testing.T) {
	a := accessory.NewOutlet(accessory.Info{Name: "ABC"})

	s, err := NewServer(NewMemStore(), a.A)
	if err != nil {
		t.Fatal(err)
	}

	body := fmt.Sprintf("{\"characteristics\":[{\"aid\":%d,\"iid\":%d,\"value\":true}]}", a.Id, a.Outlet.On.Id)
	req := httptest.NewRequest(http.MethodPut, "/characteristics", bytes.NewBuffer([]byte(body)))
	w := httptest.NewRecorder()

	s.setSession(req.RemoteAddr, &session{})

	a.Outlet.On.SetValueRequestFunc = func(v interface{}, r *http.Request) (interface{}, int) {
		return nil, JsonStatusResourceBusy
	}

	s.ss.Handler.ServeHTTP(w, req)

	r := w.Result()
	if is, want := r.StatusCode, http.StatusMultiStatus; is != want {
		t.Fatalf("%v != %v", is, want)
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}

	body = fmt.Sprintf("{\"characteristics\":[{\"aid\":%d,\"iid\":%d,\"status\":-70403}]}", a.Id, a.Outlet.On.Id)
	if is, want := string(b), body; is != want {
		t.Fatalf("%v != %v", is, want)
	}

	if is, want := a.Outlet.On.Value(), false; is != want {
		t.Fatalf("%v != %v", is, want)
	}
}

func TestGetProgrammableSwitchEvent(t *testing.T) {
	a := accessory.New(accessory.Info{Name: "ABC"}, accessory.TypeProgrammableSwitch)
	s := service.NewStatelessProgrammableSwitch()
	c := s.ProgrammableSwitchEvent
	a.AddS(s.S)
	srv, err := NewServer(NewMemStore(), a)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/characteristics?id=%d.%d", a.Id, c.Id), nil)
	w := httptest.NewRecorder()

	srv.setSession(req.RemoteAddr, &session{})
	srv.ss.Handler.ServeHTTP(w, req)

	r := w.Result()
	if is, want := r.StatusCode, http.StatusOK; is != want {
		t.Fatalf("%v != %v", is, want)
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}

	body := fmt.Sprintf("{\"characteristics\":[{\"aid\":%d,\"iid\":%d,\"value\":null}]}", a.Id, c.Id)
	if is, want := string(b), body; is != want {
		t.Fatalf("%v != %v", is, want)
	}
}
