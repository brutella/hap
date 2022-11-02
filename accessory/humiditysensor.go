package accessory

import (
	"github.com/brutella/hap/service"
)

type HumiditySensor struct {
	*A
	HumiditySensor *service.HumiditySensor
}

// NewHumiditySensor returns a humidity which implements model.HumiditySensor.
func NewHumiditySensor(info Info) *HumiditySensor {
	a := HumiditySensor{}
	a.A = New(info, TypeSensor)

	a.HumiditySensor = service.NewHumiditySensor()
	a.AddS(a.HumiditySensor.S)

	return &a
}
