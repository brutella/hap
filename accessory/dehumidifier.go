package accessory

import (
	"github.com/brutella/hap/service"
)

type Dehumidifier struct {
	*A
	Dehumidifier *service.Dehumidifier
}

// NewDehumidifier returns an outlet accessory.
func NewDehumidifier(info Info) *Dehumidifier {
	a := Dehumidifier{}
	a.A = New(info, TypeDehumidifier)

	a.Dehumidifier = service.NewDehumidifier()
	a.AddS(a.Dehumidifier.S)

	return &a
}
