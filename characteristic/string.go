package characteristic

import (
	"net/http"
)

type String struct {
	*C
}

func NewString(typ string) *String {
	c := New()
	c.Type = typ
	c.Format = FormatString

	return &String{c}
}

// SetValue sets the value of c to v.
func (c *String) SetValue(v string) {
	c.setValue(v, nil)
}

// Value returns the value of c as string.
func (c *String) Value() string {
	return c.C.value(nil).(string)
}

// OnValueRemoteUpdate calls fn when the value was updated by a client.
func (c *String) OnValueUpdate(fn func(old, new string, r *http.Request)) {
	c.OnCValueUpdate(func(c *C, new, old interface{}, r *http.Request) {
		fn(new.(string), old.(string), r)
	})
}
