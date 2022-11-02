package service

import (
	"github.com/brutella/hap/characteristic"
)

type DimmerColorTemperatureLightbulb struct {
	*S

	On               *characteristic.On
	Brightness       *characteristic.Brightness
	ColorTemperature *characteristic.ColorTemperature
}

func NewDimmerColorTemperatureLightbulb() *DimmerColorTemperatureLightbulb {
	s := DimmerColorTemperatureLightbulb{}
	s.S = New(TypeLightbulb)

	s.On = characteristic.NewOn()
	s.AddC(s.On.C)

	s.Brightness = characteristic.NewBrightness()
	s.AddC(s.Brightness.C)

	s.ColorTemperature = characteristic.NewColorTemperature()
	s.AddC(s.ColorTemperature.C)

	return &s
}
