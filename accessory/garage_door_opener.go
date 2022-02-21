package accessory

import (
	"github.com/brutella/hap/service"
)

type GarageDoorOpener struct {
	*A

	GarageDoorOpener *service.GarageDoorOpener
}

// NewGarageDoorOpener returns a garage door opener accessory.
func NewGarageDoorOpener(info Info) *GarageDoorOpener {
	a := New(info, TypeGarageDoorOpener)

	garage := service.NewGarageDoorOpener()
	a.Ss = append(a.Ss, garage.S)

	return &GarageDoorOpener{
		A:                a,
		GarageDoorOpener: garage,
	}
}
