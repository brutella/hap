package characteristic

import (
	"net/http"
)

type Bool struct {
	*C
}

func NewBool(typ string) *Bool {
	number := New()
	number.Type = typ
	number.Format = FormatBool

	return &Bool{number}
}

// SetValue sets the value of c to v.
func (c *Bool) SetValue(v bool) {
	c.setValue(v, nil)
}

// Value returns the value of c as bool.
func (c *Bool) Value() bool {
	v, _ := c.C.valueRequest(nil)
	if v == nil {
		return false
	}

	return v.(bool)
}

// OnValueRemoteUpdate calls fn when the value of the characteristic was updated.
// If the provided http request is not nil, the value was updated by a client (ex. iOS device).
func (c *Bool) OnValueUpdate(fn func(old, new bool, r *http.Request)) {
	c.OnCValueUpdate(func(c *C, new, old interface{}, r *http.Request) {
		fn(new.(bool), old.(bool), r)
	})
}

// OnValueRemoteUpdate calls fn when the value of the characteristic was updated by a client.
func (c *Bool) OnValueRemoteUpdate(fn func(v bool)) {
	c.OnCValueUpdate(func(c *C, new, old interface{}, r *http.Request) {
		if r != nil {
			fn(new.(bool))
		}
	})
}
