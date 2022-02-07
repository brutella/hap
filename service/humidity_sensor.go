// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeHumiditySensor = "82"

type HumiditySensor struct {
	*S

	CurrentRelativeHumidity *characteristic.CurrentRelativeHumidity
}

func NewHumiditySensor() *HumiditySensor {
	s := HumiditySensor{}
	s.S = New(TypeHumiditySensor)

	s.CurrentRelativeHumidity = characteristic.NewCurrentRelativeHumidity()
	s.AddC(s.CurrentRelativeHumidity.C)

	return &s
}
