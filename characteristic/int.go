package characteristic

import (
	"github.com/brutella/hap/log"

	"net/http"
)

type Int struct {
	*C
}

func NewInt(t string) *Int {
	c := New()
	c.Type = t
	return &Int{c}
}

// SetValue sets a value
func (c *Int) SetValue(v int) {
	c.setValue(v, nil)
}

func (c *Int) SetMinValue(v int) {
	c.MinVal = v
}

func (c *Int) SetMaxValue(v int) {
	c.MaxVal = v
}

func (c *Int) SetStepValue(v int) {
	c.StepVal = v
}

// Value returns the value of c as integer.
func (c *Int) Value() int {
	v, _ := c.C.valueRequest(nil)
	if v == nil {
		return 0
	}

	return v.(int)
}

func (c *Int) MinValue() int {
	return c.MinVal.(int)
}

func (c *Int) MaxValue() int {
	return c.MaxVal.(int)
}

func (c *Int) StepValue() int {
	return c.StepVal.(int)
}

// OnSetRemoteValue set c.SetValueRequestFunc and calls fn only
// if the value is going to be updated from a request.
func (c *Int) OnSetRemoteValue(fn func(v int) error) {
	c.SetValueRequestFunc = func(v interface{}, r *http.Request) int {
		if r == nil {
			return 0
		}

		if err := fn(v.(int)); err != nil {
			log.Debug.Println(err)
			return -70402
		}
		return 0
	}
}

// OnValueRemoteUpdate calls fn when the value of the characteristic was updated.
// If the provided http request is not nil, the value was updated by a client (ex. iOS device).
func (c *Int) OnValueUpdate(fn func(new, old int, r *http.Request)) {
	c.OnCValueUpdate(func(c *C, new, old interface{}, r *http.Request) {
		fn(new.(int), old.(int), r)
	})
}

// OnValueRemoteUpdate calls fn when the value of the characteristic was updated by a client.
func (c *Int) OnValueRemoteUpdate(fn func(v int)) {
	c.OnCValueUpdate(func(c *C, new, old interface{}, r *http.Request) {
		if r != nil {
			fn(new.(int))
		}
	})
}
