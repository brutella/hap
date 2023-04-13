// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeCarbonMonoxideSensor = "7F"

type CarbonMonoxideSensor struct {
	*S

	CarbonMonoxideDetected *characteristic.CarbonMonoxideDetected
	CarbonDioxideLevel     *characteristic.CarbonDioxideLevel
}

func NewCarbonMonoxideSensor() *CarbonMonoxideSensor {
	s := CarbonMonoxideSensor{}
	s.S = New(TypeCarbonMonoxideSensor)

	s.CarbonMonoxideDetected = characteristic.NewCarbonMonoxideDetected()
	s.AddC(s.CarbonMonoxideDetected.C)

	s.CarbonDioxideLevel = characteristic.NewCarbonDioxideLevel()
	s.AddC(s.CarbonDioxideLevel.C)

	return &s
}
