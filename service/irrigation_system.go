// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeIrrigationSystem = "CF"

type IrrigationSystem struct {
	*S

	Active      *characteristic.Active
	ProgramMode *characteristic.ProgramMode
	InUse       *characteristic.InUse
}

func NewIrrigationSystem() *IrrigationSystem {
	s := IrrigationSystem{}
	s.S = New(TypeIrrigationSystem)

	s.Active = characteristic.NewActive()
	s.AddC(s.Active.C)

	s.ProgramMode = characteristic.NewProgramMode()
	s.AddC(s.ProgramMode.C)

	s.InUse = characteristic.NewInUse()
	s.AddC(s.InUse.C)

	return &s
}
