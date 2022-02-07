package characteristic

import (
	"net/http"
)

type Float struct {
	*C
}

func NewFloat(t string) *Float {
	c := New()
	c.Type = t
	return &Float{c}
}

// SetValue sets a value
func (c *Float) SetValue(v float64) {
	c.setValue(v, nil)
}

func (c *Float) SetMinValue(v float64) {
	c.MinVal = v
}

func (c *Float) SetMaxValue(v float64) {
	c.MaxVal = v
}

func (c *Float) SetStepValue(v float64) {
	c.StepVal = v
}

// Value returns the value as float
func (c *Float) Value() float64 {
	return c.C.value(nil).(float64)
}

func (c *Float) MinValue() float64 {
	return c.MinVal.(float64)
}

func (c *Float) MaxValue() float64 {
	return c.MaxVal.(float64)
}

func (c *Float) StepValue() float64 {
	return c.StepVal.(float64)
}

// OnValueUpdate calls fn when the value was updated by a client.
func (c *Float) OnValueUpdate(fn func(new, old float64, r *http.Request)) {
	c.OnCValueUpdate(func(c *C, new, old interface{}, r *http.Request) {
		fn(new.(float64), old.(float64), r)
	})
}

func (c *Float) OnValueRemoteUpdate(fn func(v float64)) {
	c.OnCValueUpdate(func(c *C, new, old interface{}, r *http.Request) {
		if r != nil {
			fn(new.(float64))
		}
	})
}
