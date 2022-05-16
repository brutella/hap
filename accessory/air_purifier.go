package accessory

import (
	"github.com/brutella/hap/service"
)

type AirPurifier struct {
	*A
	AirPurifier *service.AirPurifier
}

// NewAirPurifier returns a air purifier accessory.
func NewAirPurifier(info Info) *AirPurifier {
	a := AirPurifier{}
	a.A = New(info, TypeAirPurifier)

	a.AirPurifier = service.NewAirPurifier()
	a.AddS(a.AirPurifier.S)

	return &a
}
