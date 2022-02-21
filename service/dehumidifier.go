package service

import (
	"github.com/brutella/hap/characteristic"
)

type Dehumidifier struct {
	*S

	CurrentRelativeHumidity               *characteristic.CurrentRelativeHumidity
	CurrentHumidifierDehumidifierState    *characteristic.CurrentHumidifierDehumidifierState
	TargetHumidifierDehumidifierState     *characteristic.TargetHumidifierDehumidifierState
	Active                                *characteristic.Active
	RelativeHumidityDehumidifierThreshold *characteristic.RelativeHumidityDehumidifierThreshold
}

func NewDehumidifier() *Dehumidifier {
	s := Dehumidifier{}
	s.S = New(TypeHumidifierDehumidifier)

	s.CurrentRelativeHumidity = characteristic.NewCurrentRelativeHumidity()
	s.AddC(s.CurrentRelativeHumidity.C)

	s.CurrentHumidifierDehumidifierState = characteristic.NewCurrentHumidifierDehumidifierState()
	s.AddC(s.CurrentHumidifierDehumidifierState.C)

	s.TargetHumidifierDehumidifierState = characteristic.NewTargetHumidifierDehumidifierState()
	s.AddC(s.TargetHumidifierDehumidifierState.C)

	s.Active = characteristic.NewActive()
	s.AddC(s.Active.C)

	s.RelativeHumidityDehumidifierThreshold = characteristic.NewRelativeHumidityDehumidifierThreshold()
	s.AddC(s.RelativeHumidityDehumidifierThreshold.C)

	return &s
}
