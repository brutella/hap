package service

import (
	"github.com/brutella/hap/characteristic"
)

type Cooler struct {
	*S

	Active                      *characteristic.Active
	CurrentHeaterCoolerState    *characteristic.CurrentHeaterCoolerState
	TargetHeaterCoolerState     *characteristic.TargetHeaterCoolerState
	CurrentTemperature          *characteristic.CurrentTemperature
	CoolingThresholdTemperature *characteristic.CoolingThresholdTemperature
}

func NewCooler() *Cooler {
	s := Cooler{}
	s.S = New(TypeHeaterCooler)

	s.Active = characteristic.NewActive()
	s.AddC(s.Active.C)

	s.CurrentHeaterCoolerState = characteristic.NewCurrentHeaterCoolerState()
	s.CurrentHeaterCoolerState.ValidVals = []int{
		characteristic.CurrentHeaterCoolerStateInactive,
		characteristic.CurrentHeaterCoolerStateIdle,
		characteristic.CurrentHeaterCoolerStateCooling,
	}
	s.AddC(s.CurrentHeaterCoolerState.C)

	s.TargetHeaterCoolerState = characteristic.NewTargetHeaterCoolerState()
	s.TargetHeaterCoolerState.ValidVals = []int{
		characteristic.TargetHeaterCoolerStateAuto,
		characteristic.TargetHeaterCoolerStateCool,
	}
	s.AddC(s.TargetHeaterCoolerState.C)

	s.CurrentTemperature = characteristic.NewCurrentTemperature()
	s.AddC(s.CurrentTemperature.C)

	s.CoolingThresholdTemperature = characteristic.NewCoolingThresholdTemperature()
	s.AddC(s.CoolingThresholdTemperature.C)

	return &s
}
