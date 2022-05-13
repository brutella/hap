package accessory

import (
	"github.com/brutella/hap/service"
)

type Thermostat struct {
	*A
	Thermostat *service.Thermostat
}

// NewThermostat returns a Thermostat accessory.
func NewThermostat(info Info) *Thermostat {
	a := Thermostat{}
	a.A = New(info, TypeThermostat)

	a.Thermostat = service.NewThermostat()
	a.AddS(a.Thermostat.S)

	return &a
}
