package accessory

import (
	"github.com/brutella/hap/service"
)

type WindowCovering struct {
	*A
	WindowCovering *service.WindowCovering
}

// NewWindowCovering returns a window accessory.
func NewWindowCovering(info Info) *WindowCovering {
	a := WindowCovering{}
	a.A = New(info, TypeWindowCovering)

	a.WindowCovering = service.NewWindowCovering()
	a.AddS(a.WindowCovering.S)

	return &a
}
