package characteristic

import (
	"github.com/brutella/hap/log"

	"net/http"
)

type Bool struct {
	*C
}

func NewBool(typ string) *Bool {
	c := New()
	c.Type = typ
	c.Format = FormatBool

	return &Bool{c}
}

// SetValue sets the value of c to v.
func (c *Bool) SetValue(v bool) {
	c.setValue(v, nil)
}

// Value returns the value of c as bool.
func (c *Bool) Value() bool {
	return c.C.Value().(bool)
}

// OnSetRemoteValue set c.SetValueRequestFunc and calls fn.
// If the function returns an error, the code -70402 is
// included in the HTTP response.
func (c *Bool) OnSetRemoteValue(fn func(v bool) error) {
	c.SetValueRequestFunc = func(v interface{}, r *http.Request) (interface{}, int) {
		if err := fn(v.(bool)); err != nil {
			log.Debug.Println(err)
			return nil, -70402
		}
		return nil, 0
	}
}

// OnValueRemoteUpdate calls fn when the value of the characteristic was updated.
// If the provided http request is not nil, the value was updated by a paired controller (ex. iOS device).
func (c *Bool) OnValueUpdate(fn func(old, new bool, req *http.Request)) {
	c.OnCValueUpdate(func(c *C, new, old interface{}, r *http.Request) {
		fn(new.(bool), old.(bool), r)
	})
}

// OnValueRemoteUpdate calls fn when the value of the C was updated by a paired controller (ex. iOS device).
func (c *Bool) OnValueRemoteUpdate(fn func(v bool)) {
	c.OnCValueUpdate(func(c *C, new, old interface{}, r *http.Request) {
		if r != nil {
			fn(new.(bool))
		}
	})
}
