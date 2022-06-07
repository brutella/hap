package accessory

import (
	"github.com/brutella/hap/service"
)

type Thermometer struct {
	*A
	TempSensor *service.TemperatureSensor
}

// NewTemperatureSensor returns a Thermometer which implements model.Thermometer.
func NewTemperatureSensor(info Info) *Thermometer {
	a := Thermometer{}
	a.A = New(info, TypeThermostat)

	a.TempSensor = service.NewTemperatureSensor()
	a.AddS(a.TempSensor.S)

	return &a
}
