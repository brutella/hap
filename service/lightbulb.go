// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeLightbulb = "43"

type Lightbulb struct {
	*S

	On *characteristic.On
}

func NewLightbulb() *Lightbulb {
	s := Lightbulb{}
	s.S = New(TypeLightbulb)

	s.On = characteristic.NewOn()
	s.AddC(s.On.C)

	return &s
}
