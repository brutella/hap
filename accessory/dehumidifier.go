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
	a := New(info, TypeDehumidifier)

	d := service.NewDehumidifier()
	a.Ss = append(a.Ss, d.S)

	return &Dehumidifier{
		A:            a,
		Dehumidifier: d,
	}
}
