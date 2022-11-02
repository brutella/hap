package accessory

import (
	"github.com/brutella/hap/service"
)

type OccupancySensor struct {
	*A
	OccupancySensor *service.OccupancySensor
}

// NewOccupancySensor returns a OccupancySensor which implements model.OccupancySensor.
func NewOccupancySensor(info Info) *OccupancySensor {
	a := OccupancySensor{}
	a.A = New(info, TypeSensor)

	a.OccupancySensor = service.NewOccupancySensor()
	a.AddS(a.OccupancySensor.S)

	return &a
}
