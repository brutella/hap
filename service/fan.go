// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeFan = "40"

type Fan struct {
	*S

	On *characteristic.On
}

func NewFan() *Fan {
	s := Fan{}
	s.S = New(TypeFan)

	s.On = characteristic.NewOn()
	s.AddC(s.On.C)

	return &s
}
