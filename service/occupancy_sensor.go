// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeOccupancySensor = "86"

type OccupancySensor struct {
	*S

	OccupancyDetected *characteristic.OccupancyDetected
}

func NewOccupancySensor() *OccupancySensor {
	s := OccupancySensor{}
	s.S = New(TypeOccupancySensor)

	s.OccupancyDetected = characteristic.NewOccupancyDetected()
	s.AddC(s.OccupancyDetected.C)

	return &s
}
