// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeHumidifierDehumidifier = "BD"

type HumidifierDehumidifier struct {
	*S

	CurrentRelativeHumidity            *characteristic.CurrentRelativeHumidity
	CurrentHumidifierDehumidifierState *characteristic.CurrentHumidifierDehumidifierState
	TargetHumidifierDehumidifierState  *characteristic.TargetHumidifierDehumidifierState
	Active                             *characteristic.Active
}

func NewHumidifierDehumidifier() *HumidifierDehumidifier {
	s := HumidifierDehumidifier{}
	s.S = New(TypeHumidifierDehumidifier)

	s.CurrentRelativeHumidity = characteristic.NewCurrentRelativeHumidity()
	s.AddC(s.CurrentRelativeHumidity.C)

	s.CurrentHumidifierDehumidifierState = characteristic.NewCurrentHumidifierDehumidifierState()
	s.AddC(s.CurrentHumidifierDehumidifierState.C)

	s.TargetHumidifierDehumidifierState = characteristic.NewTargetHumidifierDehumidifierState()
	s.AddC(s.TargetHumidifierDehumidifierState.C)

	s.Active = characteristic.NewActive()
	s.AddC(s.Active.C)

	return &s
}
