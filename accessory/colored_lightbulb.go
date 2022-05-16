package accessory

import (
	"github.com/brutella/hap/service"
)

type ColoredLightbulb struct {
	*A
	Lightbulb *service.ColoredLightbulb
}

// NewLightbulb returns an light bulb accessory.
func NewColoredLightbulb(info Info) *ColoredLightbulb {
	a := ColoredLightbulb{}
	a.A = New(info, TypeLightbulb)

	a.Lightbulb = service.NewColoredLightbulb()
	a.AddS(a.Lightbulb.S)

	return &a
}
