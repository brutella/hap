package accessory

import (
	"github.com/brutella/hap/service"
)

type Lightbulb struct {
	*A
	Lightbulb *service.Lightbulb
}

// NewLightbulb returns an light bulb accessory.
func NewLightbulb(info Info) *Lightbulb {
	a := New(info, TypeLightbulb)

	l := service.NewLightbulb()
	a.AddS(l.S)

	return &Lightbulb{
		A:         a,
		Lightbulb: l,
	}
}
