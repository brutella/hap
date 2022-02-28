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
	v, _ := c.C.valueRequest(nil)
	if v == nil {
		return []byte{}
	}

	str := v.(string)
	if b, err := base64.StdEncoding.DecodeString(str); err != nil {
		return []byte{}
	} else {
		return b
	}
}

// OnSetRemoteValue set c.SetValueRequestFunc and calls fn only
// if the value is going to be updated from a request.
func (c *Bytes) OnSetRemoteValue(fn func(v []byte) error) {
	c.SetValueRequestFunc = func(v interface{}, r *http.Request) int {
		if r == nil {
			return 0
		}

		str, _ := base64.StdEncoding.DecodeString(v.(string))

		if err := fn(str); err != nil {
			log.Debug.Println(err)
			return -70402
		}
		return 0
	}
}

func (c *Bytes) OnValueUpdate(fn func(old, new []byte, r *http.Request)) {
	c.OnCValueUpdate(func(c *C, new, old interface{}, r *http.Request) {
		newVal, _ := base64.StdEncoding.DecodeString(new.(string))
		oldVal, _ := base64.StdEncoding.DecodeString(old.(string))
		fn(newVal, oldVal, r)
	})
}

func (c *Bytes) OnValueRemoteUpdate(fn func(v []byte)) {
	c.OnValueUpdate(func(new, old []byte, r *http.Request) {
		if r != nil {
			fn(new)
		}
	})
}

func base64FromBytes(v []byte) string {
	return base64.StdEncoding.EncodeToString(v)
}
