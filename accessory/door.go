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
	a := New(info, TypeDoor)

	door := service.NewDoor()
	a.Ss = append(a.Ss, door.S)

	return &Door{
		A:    a,
		Door: door,
	}
}
