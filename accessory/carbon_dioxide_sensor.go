package accessory

import (
	"github.com/brutella/hap/service"
)

type CarbonDioxideSensor struct {
	*A
	CarbonDioxideSensor *service.CarbonDioxideSensor
}

// NewCarbonDioxideSensor returns a CarbonDioxideSensor which implements model.CarbonDioxideSensor.
func NewCarbonDioxideSensor(info Info) *CarbonDioxideSensor {
	a := CarbonDioxideSensor{}
	a.A = New(info, TypeSensor)

	a.CarbonDioxideSensor = service.NewCarbonDioxideSensor()
	a.AddS(a.CarbonDioxideSensor.S)

	return &a
}
