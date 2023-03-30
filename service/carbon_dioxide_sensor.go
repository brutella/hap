// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeCarbonDioxideSensor = "97"

type CarbonDioxideSensor struct {
	*S

	CarbonDioxideDetected *characteristic.CarbonDioxideDetected
}

func NewCarbonDioxideSensor() *CarbonDioxideSensor {
	s := CarbonDioxideSensor{}
	s.S = New(TypeCarbonDioxideSensor)

	s.CarbonDioxideDetected = characteristic.NewCarbonDioxideDetected()
	s.AddC(s.CarbonDioxideDetected.C)

	return &s
}
