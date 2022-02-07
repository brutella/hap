// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeGarageDoorOpener = "41"

type GarageDoorOpener struct {
	*S

	CurrentDoorState    *characteristic.CurrentDoorState
	TargetDoorState     *characteristic.TargetDoorState
	ObstructionDetected *characteristic.ObstructionDetected
}

func NewGarageDoorOpener() *GarageDoorOpener {
	s := GarageDoorOpener{}
	s.S = New(TypeGarageDoorOpener)

	s.CurrentDoorState = characteristic.NewCurrentDoorState()
	s.AddC(s.CurrentDoorState.C)

	s.TargetDoorState = characteristic.NewTargetDoorState()
	s.AddC(s.TargetDoorState.C)

	s.ObstructionDetected = characteristic.NewObstructionDetected()
	s.AddC(s.ObstructionDetected.C)

	return &s
}
