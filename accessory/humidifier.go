package accessory

import (
	"github.com/brutella/hap/service"
)

type Humidifier struct {
	*A

	Humidifier *service.Humidifier
}

// NewHumidifier returns an outlet accessory.
func NewHumidifier(info Info) *Humidifier {
	a := New(info, TypeHumidifier)

	h := service.NewHumidifier()
	a.Ss = append(a.Ss, h.S)

	return &Humidifier{
		A:          a,
		Humidifier: h,
	}
}
