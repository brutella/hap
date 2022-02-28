package characteristic

import (
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
