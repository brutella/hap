// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeDoor = "81"

type Door struct {
	*S

	CurrentPosition *characteristic.CurrentPosition
	PositionState   *characteristic.PositionState
	TargetPosition  *characteristic.TargetPosition
}

func NewDoor() *Door {
	s := Door{}
	s.S = New(TypeDoor)

	s.CurrentPosition = characteristic.NewCurrentPosition()
	s.AddC(s.CurrentPosition.C)

	s.PositionState = characteristic.NewPositionState()
	s.AddC(s.PositionState.C)

	s.TargetPosition = characteristic.NewTargetPosition()
	s.AddC(s.TargetPosition.C)

	return &s
}
