package accessory

import (
	"github.com/brutella/hap/service"
)

type Fan struct {
	*A

	Fan *service.Fan
}

// NewFan returns a fan accessory.
func NewFan(info Info) *Fan {
	a := New(info, TypeFan)

	fan := service.NewFan()
	a.Ss = append(a.Ss, fan.S)

	return &Fan{
		A:   a,
		Fan: fan,
	}
}
