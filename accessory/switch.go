package accessory

import (
	"github.com/brutella/hap/service"
)

type Switch struct {
	*A
	Switch *service.Switch
}

// NewSwitch returns a switch which implements model.Switch.
func NewSwitch(info Info) *Switch {
	a := Switch{}
	a.A = New(info, TypeSwitch)
	
	a.Switch = service.NewSwitch()
	a.AddS(a.Switch.S)

	return &a
}
