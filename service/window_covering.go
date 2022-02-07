// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeWindowCovering = "8C"

type WindowCovering struct {
	*S

	CurrentPosition *characteristic.CurrentPosition
	TargetPosition  *characteristic.TargetPosition
	PositionState   *characteristic.PositionState
}

func NewWindowCovering() *WindowCovering {
	s := WindowCovering{}
	s.S = New(TypeWindowCovering)

	s.CurrentPosition = characteristic.NewCurrentPosition()
	s.AddC(s.CurrentPosition.C)

	s.TargetPosition = characteristic.NewTargetPosition()
	s.AddC(s.TargetPosition.C)

	s.PositionState = characteristic.NewPositionState()
	s.AddC(s.PositionState.C)

	return &s
}
