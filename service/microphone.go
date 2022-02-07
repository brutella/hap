// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeMicrophone = "112"

type Microphone struct {
	*S

	Volume *characteristic.Volume
	Mute   *characteristic.Mute
}

func NewMicrophone() *Microphone {
	s := Microphone{}
	s.S = New(TypeMicrophone)

	s.Volume = characteristic.NewVolume()
	s.AddC(s.Volume.C)

	s.Mute = characteristic.NewMute()
	s.AddC(s.Mute.C)

	return &s
}
