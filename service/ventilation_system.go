// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

type VentilationSystem struct {
	*S

	Active                  *characteristic.Active
	CurrentAirPurifierState *characteristic.CurrentAirPurifierState
	TargetAirPurifierState  *characteristic.TargetAirPurifierState
	CarbonDioxideLevel      *characteristic.CarbonDioxideLevel
	PM2_5Density            *characteristic.PM2_5Density
	CurrentTemperature      *characteristic.CurrentTemperature
	CurrentRelativeHumidity *characteristic.CurrentRelativeHumidity
}

func NewVentilationSystem() *VentilationSystem {
	s := VentilationSystem{}
	s.S = New(TypeAirPurifier)

	s.Active = characteristic.NewActive()
	s.AddC(s.Active.C)

	s.CurrentAirPurifierState = characteristic.NewCurrentAirPurifierState()
	s.AddC(s.CurrentAirPurifierState.C)

	s.TargetAirPurifierState = characteristic.NewTargetAirPurifierState()
	s.AddC(s.TargetAirPurifierState.C)

	s.CarbonDioxideLevel = characteristic.NewCarbonDioxideLevel()
	s.AddC(s.CarbonDioxideLevel.C)

	s.PM2_5Density = characteristic.NewPM2_5Density()
	s.AddC(s.PM2_5Density.C)

	s.CurrentTemperature = characteristic.NewCurrentTemperature()
	s.AddC(s.CurrentTemperature.C)

	s.CurrentRelativeHumidity = characteristic.NewCurrentRelativeHumidity()
	s.AddC(s.CurrentRelativeHumidity.C)

	return &s
}
