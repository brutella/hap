// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeTemperatureSensor = "8A"

type TemperatureSensor struct {
	*S

	CurrentTemperature *characteristic.CurrentTemperature
}

func NewTemperatureSensor() *TemperatureSensor {
	s := TemperatureSensor{}
	s.S = New(TypeTemperatureSensor)

	s.CurrentTemperature = characteristic.NewCurrentTemperature()
	s.AddC(s.CurrentTemperature.C)

	return &s
}
