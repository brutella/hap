package service

import (
	"github.com/brutella/hap/characteristic"
)

type DimmerLightbulb struct {
	*S

	On         *characteristic.On
	Brightness *characteristic.Brightness
}

func NewDimmerLightbulb() *DimmerLightbulb {
	s := DimmerLightbulb{}
	s.S = New(TypeLightbulb)

	s.On = characteristic.NewOn()
	s.AddC(s.On.C)

	s.Brightness = characteristic.NewBrightness()
	s.AddC(s.Brightness.C)

	return &s
}
