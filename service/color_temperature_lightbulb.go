package service

import (
	"github.com/brutella/hap/characteristic"
)

type ColorTemperatureLightbulb struct {
	*S

	On               *characteristic.On
	ColorTemperature *characteristic.ColorTemperature
}

func NewColorTemperatureLightbulb() *ColorTemperatureLightbulb {
	s := ColorTemperatureLightbulb{}
	s.S = New(TypeLightbulb)

	s.On = characteristic.NewOn()
	s.AddC(s.On.C)

	s.ColorTemperature = characteristic.NewColorTemperature()
	s.AddC(s.ColorTemperature.C)

	return &s
}
