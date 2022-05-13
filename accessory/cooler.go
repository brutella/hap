package accessory

import (
	"github.com/brutella/hap/service"
)

type Cooler struct {
	*A
	Cooler *service.Cooler
}

// NewCooler returns a cooler accessory.
func NewCooler(info Info) *Cooler {
	a := Cooler{}
	a.A = New(info, TypeAirConditioner)

	a.Cooler = service.NewCooler()
	a.AddS(a.Cooler.S)

	return &a
}
