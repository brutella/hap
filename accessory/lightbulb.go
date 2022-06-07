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
	a := Lightbulb{}
	a.A = New(info, TypeLightbulb)

	a.Lightbulb = service.NewLightbulb()
	a.AddS(a.Lightbulb.S)

	return &a
}
