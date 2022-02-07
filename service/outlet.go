// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeOutlet = "47"

type Outlet struct {
	*S

	On          *characteristic.On
	OutletInUse *characteristic.OutletInUse
}

func NewOutlet() *Outlet {
	s := Outlet{}
	s.S = New(TypeOutlet)

	s.On = characteristic.NewOn()
	s.AddC(s.On.C)

	s.OutletInUse = characteristic.NewOutletInUse()
	s.AddC(s.OutletInUse.C)

	return &s
}
