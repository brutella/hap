package characteristic

import (
	"github.com/brutella/hap/log"

	"encoding/base64"
	"net/http"
)

type Bytes struct {
	*C
}

func NewBytes(t string) *Bytes {
	s := New()
	s.Type = t
	s.Format = FormatTLV8

	return &Bytes{s}
}

func (c *Bytes) SetValue(v []byte) {
	c.setValue(base64FromBytes(v), nil)
}

// Value returns the value of c as byte array.
func (c *Bytes) Value() []byte {
	str := c.C.Value().(string)
	if b, err := base64.StdEncoding.DecodeString(str); err != nil {
		return []byte{}
	} else {
		return b
	}
}

// OnSetRemoteValue set c.SetValueRequestFunc and calls fn.
// If the function returns an error, the code -70402 is
// included in the HTTP response.
func (c *Bytes) OnSetRemoteValue(fn func(v []byte) error) {
	c.SetValueRequestFunc = func(v interface{}, r *http.Request) (interface{}, int) {
		str, _ := base64.StdEncoding.DecodeString(v.(string))
		if err := fn(str); err != nil {
			log.Debug.Println(err)
			return nil, -70402
		}
		return nil, 0
	}
}

// OnValueRemoteUpdate calls fn when the value of the characteristic was updated.
// If the provided http request is not nil, the value was updated by a paired controller (ex. iOS device).
func (c *Bytes) OnValueRemoteUpdate(fn func(v []byte)) {
	c.OnValueUpdate(func(new, old []byte, r *http.Request) {
		if r != nil {
			fn(new)
		}
	})
}

// OnValueRemoteUpdate calls fn when the value of the C was updated by a paired controller (ex. iOS device).
func (c *Bytes) OnValueUpdate(fn func(old, new []byte, r *http.Request)) {
	c.OnCValueUpdate(func(c *C, new, old interface{}, r *http.Request) {
		newVal, _ := base64.StdEncoding.DecodeString(new.(string))
		oldVal, _ := base64.StdEncoding.DecodeString(old.(string))
		fn(newVal, oldVal, r)
	})
}

func base64FromBytes(v []byte) string {
	return base64.StdEncoding.EncodeToString(v)
}
