// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeWindow = "8B"

type Window struct {
	*S

	CurrentPosition *characteristic.CurrentPosition
	TargetPosition  *characteristic.TargetPosition
	PositionState   *characteristic.PositionState
}

func NewWindow() *Window {
	s := Window{}
	s.S = New(TypeWindow)

	s.CurrentPosition = characteristic.NewCurrentPosition()
	s.AddC(s.CurrentPosition.C)

	s.TargetPosition = characteristic.NewTargetPosition()
	s.AddC(s.TargetPosition.C)

	s.PositionState = characteristic.NewPositionState()
	s.AddC(s.PositionState.C)

	return &s
}
