// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeLeakSensor = "83"

type LeakSensor struct {
	*S

	LeakDetected *characteristic.LeakDetected
}

func NewLeakSensor() *LeakSensor {
	s := LeakSensor{}
	s.S = New(TypeLeakSensor)

	s.LeakDetected = characteristic.NewLeakDetected()
	s.AddC(s.LeakDetected.C)

	return &s
}
