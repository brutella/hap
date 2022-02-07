// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeAirQualitySensor = "8D"

type AirQualitySensor struct {
	*S

	AirQuality *characteristic.AirQuality
}

func NewAirQualitySensor() *AirQualitySensor {
	s := AirQualitySensor{}
	s.S = New(TypeAirQualitySensor)

	s.AirQuality = characteristic.NewAirQuality()
	s.AddC(s.AirQuality.C)

	return &s
}
