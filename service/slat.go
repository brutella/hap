// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeSlat = "B9"

type Slat struct {
	*S

	SlatType         *characteristic.SlatType
	CurrentSlatState *characteristic.CurrentSlatState
}

func NewSlat() *Slat {
	s := Slat{}
	s.S = New(TypeSlat)

	s.SlatType = characteristic.NewSlatType()
	s.AddC(s.SlatType.C)

	s.CurrentSlatState = characteristic.NewCurrentSlatState()
	s.AddC(s.CurrentSlatState.C)

	return &s
}
