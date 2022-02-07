// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeDoorbell = "121"

type Doorbell struct {
	*S

	ProgrammableSwitchEvent *characteristic.ProgrammableSwitchEvent
}

func NewDoorbell() *Doorbell {
	s := Doorbell{}
	s.S = New(TypeDoorbell)

	s.ProgrammableSwitchEvent = characteristic.NewProgrammableSwitchEvent()
	s.AddC(s.ProgrammableSwitchEvent.C)

	return &s
}
