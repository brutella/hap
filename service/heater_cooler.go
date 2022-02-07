// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeHeaterCooler = "BC"

type HeaterCooler struct {
	*S

	Active                   *characteristic.Active
	CurrentHeaterCoolerState *characteristic.CurrentHeaterCoolerState
	TargetHeaterCoolerState  *characteristic.TargetHeaterCoolerState
	CurrentTemperature       *characteristic.CurrentTemperature
}

func NewHeaterCooler() *HeaterCooler {
	s := HeaterCooler{}
	s.S = New(TypeHeaterCooler)

	s.Active = characteristic.NewActive()
	s.AddC(s.Active.C)

	s.CurrentHeaterCoolerState = characteristic.NewCurrentHeaterCoolerState()
	s.AddC(s.CurrentHeaterCoolerState.C)

	s.TargetHeaterCoolerState = characteristic.NewTargetHeaterCoolerState()
	s.AddC(s.TargetHeaterCoolerState.C)

	s.CurrentTemperature = characteristic.NewCurrentTemperature()
	s.AddC(s.CurrentTemperature.C)

	return &s
}
