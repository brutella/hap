// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeLightSensor = "84"

type LightSensor struct {
	*S

	CurrentAmbientLightLevel *characteristic.CurrentAmbientLightLevel
}

func NewLightSensor() *LightSensor {
	s := LightSensor{}
	s.S = New(TypeLightSensor)

	s.CurrentAmbientLightLevel = characteristic.NewCurrentAmbientLightLevel()
	s.AddC(s.CurrentAmbientLightLevel.C)

	return &s
}
