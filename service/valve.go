// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeValve = "D0"

type Valve struct {
	*S

	Active    *characteristic.Active
	InUse     *characteristic.InUse
	ValveType *characteristic.ValveType
}

func NewValve() *Valve {
	s := Valve{}
	s.S = New(TypeValve)

	s.Active = characteristic.NewActive()
	s.AddC(s.Active.C)

	s.InUse = characteristic.NewInUse()
	s.AddC(s.InUse.C)

	s.ValveType = characteristic.NewValveType()
	s.AddC(s.ValveType.C)

	return &s
}
