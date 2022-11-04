package accessory

import (
	"github.com/brutella/hap/service"
)

type Heater_New struct {
	*A
	Heater *service.Heater_New
}

// NewHeater_New returns a heater_new accessory.
func NewHeater_New(info Info) *Heater_New {
	a := Heater_New{}
	a.A = New(info, TypeHeater)

	a.Heater = service.NewHeater_New()
	a.AddS(a.Heater.S)

	return &a
}
