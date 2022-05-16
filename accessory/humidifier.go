package accessory

import (
	"github.com/brutella/hap/service"
)

type Humidifier struct {
	*A
	Humidifier *service.Humidifier
}

// NewHumidifier returns an outlet accessory.
func NewHumidifier(info Info) *Humidifier {
	a := Humidifier{}
	a.A = New(info, TypeHumidifier)

	a.Humidifier = service.NewHumidifier()
	a.AddS(a.Humidifier.S)

	return &a
}
