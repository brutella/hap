package service

import (
	"github.com/brutella/hap/characteristic"
)

type Humidifier struct {
	*S

	CurrentRelativeHumidity             *characteristic.CurrentRelativeHumidity
	CurrentHumidifierDehumidifierState  *characteristic.CurrentHumidifierDehumidifierState
	TargetHumidifierDehumidifierState   *characteristic.TargetHumidifierDehumidifierState
	Active                              *characteristic.Active
	RelativeHumidityHumidifierThreshold *characteristic.RelativeHumidityHumidifierThreshold
}

func NewHumidifier() *Humidifier {
	s := Humidifier{}
	s.S = New(TypeHumidifierDehumidifier)

	s.CurrentRelativeHumidity = characteristic.NewCurrentRelativeHumidity()
	s.AddC(s.CurrentRelativeHumidity.C)

	s.CurrentHumidifierDehumidifierState = characteristic.NewCurrentHumidifierDehumidifierState()
	s.CurrentHumidifierDehumidifierState.ValidVals = []int{
		characteristic.CurrentHumidifierDehumidifierStateInactive,
		characteristic.CurrentHumidifierDehumidifierStateIdle,
		characteristic.CurrentHumidifierDehumidifierStateHumidifying,
	}
	s.AddC(s.CurrentHumidifierDehumidifierState.C)

	s.TargetHumidifierDehumidifierState = characteristic.NewTargetHumidifierDehumidifierState()
	s.TargetHumidifierDehumidifierState.ValidVals = []int{
		characteristic.TargetHumidifierDehumidifierStateHumidifier,
	}
	s.TargetHumidifierDehumidifierState.SetValue(characteristic.TargetHumidifierDehumidifierStateHumidifier)
	s.AddC(s.TargetHumidifierDehumidifierState.C)

	s.Active = characteristic.NewActive()
	s.AddC(s.Active.C)

	s.RelativeHumidityHumidifierThreshold = characteristic.NewRelativeHumidityHumidifierThreshold()
	s.AddC(s.RelativeHumidityHumidifierThreshold.C)

	return &s
}
