// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeBatteryService = "96"

type BatteryService struct {
	*S

	BatteryLevel     *characteristic.BatteryLevel
	ChargingState    *characteristic.ChargingState
	StatusLowBattery *characteristic.StatusLowBattery
}

func NewBatteryService() *BatteryService {
	s := BatteryService{}
	s.S = New(TypeBatteryService)

	s.BatteryLevel = characteristic.NewBatteryLevel()
	s.AddC(s.BatteryLevel.C)

	s.ChargingState = characteristic.NewChargingState()
	s.AddC(s.ChargingState.C)

	s.StatusLowBattery = characteristic.NewStatusLowBattery()
	s.AddC(s.StatusLowBattery.C)

	return &s
}
