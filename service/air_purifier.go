// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeAirPurifier = "BB"

type AirPurifier struct {
	*S

	Active                  *characteristic.Active
	CurrentAirPurifierState *characteristic.CurrentAirPurifierState
	TargetAirPurifierState  *characteristic.TargetAirPurifierState
}

func NewAirPurifier() *AirPurifier {
	s := AirPurifier{}
	s.S = New(TypeAirPurifier)

	s.Active = characteristic.NewActive()
	s.AddC(s.Active.C)

	s.CurrentAirPurifierState = characteristic.NewCurrentAirPurifierState()
	s.AddC(s.CurrentAirPurifierState.C)

	s.TargetAirPurifierState = characteristic.NewTargetAirPurifierState()
	s.AddC(s.TargetAirPurifierState.C)

	return &s
}
