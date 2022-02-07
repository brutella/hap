// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeSwitch = "49"

type Switch struct {
	*S

	On *characteristic.On
}

func NewSwitch() *Switch {
	s := Switch{}
	s.S = New(TypeSwitch)

	s.On = characteristic.NewOn()
	s.AddC(s.On.C)

	return &s
}
