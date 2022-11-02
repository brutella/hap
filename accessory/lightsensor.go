package accessory

import (
	"github.com/brutella/hap/service"
)

type LightSensor struct {
	*A
	LightSensor *service.LightSensor
}

// NewLightSensor returns a LightSenor which implements model.LightSensor.
func NewLightSensor(info Info) *LightSensor {
	a := LightSensor{}
	a.A = New(info, TypeSensor)

	a.LightSensor = service.NewLightSensor()
	a.AddS(a.LightSensor.S)

	return &a
}
