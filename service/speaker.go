// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeSpeaker = "113"

type Speaker struct {
	*S

	Mute *characteristic.Mute
}

func NewSpeaker() *Speaker {
	s := Speaker{}
	s.S = New(TypeSpeaker)

	s.Mute = characteristic.NewMute()
	s.AddC(s.Mute.C)

	return &s
}
