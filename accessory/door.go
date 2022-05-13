package accessory

import (
	"github.com/brutella/hap/service"
)

type Door struct {
	*A
	Door *service.Door
}

// NewDoor returns a door accessory.
func NewDoor(info Info) *Door {
	a := Door{}
	a.A = New(info, TypeDoor)

	a.Door = service.NewDoor()
	a.AddS(a.Door.S)

	return &a
}
