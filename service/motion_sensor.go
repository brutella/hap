// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeMotionSensor = "85"

type MotionSensor struct {
	*S

	MotionDetected *characteristic.MotionDetected
}

func NewMotionSensor() *MotionSensor {
	s := MotionSensor{}
	s.S = New(TypeMotionSensor)

	s.MotionDetected = characteristic.NewMotionDetected()
	s.AddC(s.MotionDetected.C)

	return &s
}
