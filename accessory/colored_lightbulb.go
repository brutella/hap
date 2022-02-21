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
	a := New(info, TypeLightbulb)

	l := service.NewColoredLightbulb()
	a.Ss = append(a.Ss, l.S)

	return &ColoredLightbulb{
		A:         a,
		Lightbulb: l,
	}
}
