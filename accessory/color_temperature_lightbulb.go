package accessory

import (
	"github.com/brutella/hap/service"
)

type ColorTemperatureLightbulb struct {
	*A
	Lightbulb *service.ColorTemperatureLightbulb
}

// NewColorTemperatureLightbulb returns an ColorTemperatureLight bulb accessory.
func NewColorTemperatureLightbulb(info Info) *ColorTemperatureLightbulb {
	a := ColorTemperatureLightbulb{}
	a.A = New(info, TypeLightbulb)

	a.Lightbulb = service.NewColorTemperatureLightbulb()
	a.AddS(a.Lightbulb.S)

	return &a
}
