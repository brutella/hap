package accessory

import (
	"github.com/brutella/hap/service"
)

type AirQualitySensor struct {
	*A
	AirQualitySensor *service.AirQualitySensor
}

// NewAirQualitySensor returns a AirQualitySensor which implements model.AirQualitySensor.
func NewAirQualitySensor(info Info) *AirQualitySensor {
	a := AirQualitySensor{}
	a.A = New(info, TypeSensor)

	a.AirQualitySensor = service.NewAirQualitySensor()
	a.AddS(a.AirQualitySensor.S)

	return &a
}
