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
	a := GarageDoorOpener{}
	a.A = New(info, TypeGarageDoorOpener)

	a.GarageDoorOpener = service.NewGarageDoorOpener()
	a.AddS(a.GarageDoorOpener.S)

	return &a
}
