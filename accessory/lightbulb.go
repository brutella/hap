package accessory

import (
	"github.com/brutella/hap/service"
)

type Lightbulb struct {
	*A
	Lightbulb *service.Lightbulb
}

// NewLightbulb returns an light bulb accessory which one light bulb service.
func NewLightbulb(info Info) *Lightbulb {
	a := New(info, TypeLightbulb)

	l := service.NewLightbulb()
	a.Ss = append(a.Ss, l.S)

	return &Lightbulb{
		A:         a,
		Lightbulb: l,
	}
}
