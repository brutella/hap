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
	v, _ := c.C.valueRequest(nil)
	if v == nil {
		return ""
	}

	return v.(string)
}

// OnValueRemoteUpdate calls fn when the value of the characteristic was updated.
// If the provided http request is not nil, the value was updated by a client (ex. iOS device).
func (c *String) OnValueUpdate(fn func(new, old string, r *http.Request)) {
	c.OnCValueUpdate(func(c *C, new, old interface{}, r *http.Request) {
		fn(new.(string), old.(string), r)
	})
}

// OnValueRemoteUpdate calls fn when the value of the characteristic was updated by a client.
func (c *String) OnValueRemoteUpdate(fn func(v string)) {
	c.OnCValueUpdate(func(c *C, new, old interface{}, r *http.Request) {
		if r != nil {
			fn(new.(string))
		}
	})
}
