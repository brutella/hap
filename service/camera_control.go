// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeCameraControl = "111"

type CameraControl struct {
	*S

	On *characteristic.On
}

func NewCameraControl() *CameraControl {
	s := CameraControl{}
	s.S = New(TypeCameraControl)

	s.On = characteristic.NewOn()
	s.AddC(s.On.C)

	return &s
}
