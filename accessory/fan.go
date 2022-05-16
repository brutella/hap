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
	a := Fan{}
	a.A = New(info, TypeFan)

	a.Fan = service.NewFan()
	a.AddS(a.Fan.S)

	return &a
}
