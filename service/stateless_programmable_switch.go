// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeStatelessProgrammableSwitch = "89"

type StatelessProgrammableSwitch struct {
	*S

	ProgrammableSwitchEvent *characteristic.ProgrammableSwitchEvent
}

func NewStatelessProgrammableSwitch() *StatelessProgrammableSwitch {
	s := StatelessProgrammableSwitch{}
	s.S = New(TypeStatelessProgrammableSwitch)

	s.ProgrammableSwitchEvent = characteristic.NewProgrammableSwitchEvent()
	s.AddC(s.ProgrammableSwitchEvent.C)

	return &s
}
