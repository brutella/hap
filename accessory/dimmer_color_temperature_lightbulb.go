package accessory

import (
	"github.com/brutella/hap/service"
)

type DimmerColorTemperatureLightbulb struct {
	*A
	Lightbulb *service.DimmerColorTemperatureLightbulb
}

// NewDimmerColorTemperatureLightbulb returns an DimmerColorTemperatureLight bulb accessory.
func NewDimmerColorTemperatureLightbulb(info Info) *DimmerColorTemperatureLightbulb {
	a := DimmerColorTemperatureLightbulb{}
	a.A = New(info, TypeLightbulb)

	a.Lightbulb = service.NewDimmerColorTemperatureLightbulb()
	a.AddS(a.Lightbulb.S)

	return &a
}
