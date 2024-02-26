package characteristic

import (
	"github.com/brutella/hap/log"

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
	return c.C.Value().(string)
}

// OnSetRemoteValue set c.SetValueRequestFunc and calls fn.
// If the function returns an error, the code -70402 is
// included in the HTTP response.
func (c *String) OnSetRemoteValue(fn func(v string) error) {
	c.SetValueRequestFunc = func(v interface{}, r *http.Request) (interface{}, int) {
		if err := fn(v.(string)); err != nil {
			log.Debug.Println(err)
			return nil, -70402
		}
		return nil, 0
	}
}

// OnValueRemoteUpdate calls fn when the value of the characteristic was updated.
// If the provided http request is not nil, the value was updated by a paired controller (ex. iOS device).
func (c *String) OnValueUpdate(fn func(new, old string, r *http.Request)) {
	c.OnCValueUpdate(func(c *C, new, old interface{}, r *http.Request) {
		fn(new.(string), old.(string), r)
	})
}

// OnValueRemoteUpdate calls fn when the value of the C was updated by a paired controller (ex. iOS device).
func (c *String) OnValueRemoteUpdate(fn func(v string)) {
	c.OnCValueUpdate(func(c *C, new, old interface{}, r *http.Request) {
		if r != nil {
			fn(new.(string))
		}
	})
}
