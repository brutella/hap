// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeThermostat = "4A"

type Thermostat struct {
	*S

	CurrentHeatingCoolingState *characteristic.CurrentHeatingCoolingState
	TargetHeatingCoolingState  *characteristic.TargetHeatingCoolingState
	CurrentTemperature         *characteristic.CurrentTemperature
	TargetTemperature          *characteristic.TargetTemperature
	TemperatureDisplayUnits    *characteristic.TemperatureDisplayUnits
}

func NewThermostat() *Thermostat {
	s := Thermostat{}
	s.S = New(TypeThermostat)

	s.CurrentHeatingCoolingState = characteristic.NewCurrentHeatingCoolingState()
	s.AddC(s.CurrentHeatingCoolingState.C)

	s.TargetHeatingCoolingState = characteristic.NewTargetHeatingCoolingState()
	s.AddC(s.TargetHeatingCoolingState.C)

	s.CurrentTemperature = characteristic.NewCurrentTemperature()
	s.AddC(s.CurrentTemperature.C)

	s.TargetTemperature = characteristic.NewTargetTemperature()
	s.AddC(s.TargetTemperature.C)

	s.TemperatureDisplayUnits = characteristic.NewTemperatureDisplayUnits()
	s.AddC(s.TemperatureDisplayUnits.C)

	return &s
}
