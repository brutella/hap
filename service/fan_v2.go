// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeFanV2 = "B7"

type FanV2 struct {
	*S

	Active *characteristic.Active
}

func NewFanV2() *FanV2 {
	s := FanV2{}
	s.S = New(TypeFanV2)

	s.Active = characteristic.NewActive()
	s.AddC(s.Active.C)

	return &s
}
