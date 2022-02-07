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
	return c.C.value(nil).(bool)
}

func (c *Bool) OnValueUpdate(fn func(old, new bool, r *http.Request)) {
	c.OnCValueUpdate(func(c *C, new, old interface{}, r *http.Request) {
		fn(new.(bool), old.(bool), r)
	})
}

func (c *Bool) OnValueRemoteUpdate(fn func(v bool)) {
	c.OnCValueUpdate(func(c *C, new, old interface{}, r *http.Request) {
		if r != nil {
			fn(new.(bool))
		}
	})
}
