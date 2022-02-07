package characteristic

import (
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
	return c.C.value(nil).(int)
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

// OnValueRemoteUpdate calls fn when the value was updated by a client.
func (c *Int) OnValueUpdate(fn func(new, old int, r *http.Request)) {
	c.OnCValueUpdate(func(c *C, new, old interface{}, r *http.Request) {
		fn(new.(int), old.(int), r)
	})
}

func (c *Int) OnValueRemoteUpdate(fn func(v int)) {
	c.OnCValueUpdate(func(c *C, new, old interface{}, r *http.Request) {
		if r != nil {
			fn(new.(int))
		}
	})
}
