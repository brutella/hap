// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeCarbonMonoxideSensor = "7F"

type CarbonMonoxideSensor struct {
	*S

	CarbonMonoxideDetected *characteristic.CarbonMonoxideDetected
}

func NewCarbonMonoxideSensor() *CarbonMonoxideSensor {
	s := CarbonMonoxideSensor{}
	s.S = New(TypeCarbonMonoxideSensor)

	s.CarbonMonoxideDetected = characteristic.NewCarbonMonoxideDetected()
	s.AddC(s.CarbonMonoxideDetected.C)

	return &s
}
