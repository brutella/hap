package accessory

import (
	"github.com/brutella/hap/service"
)

type Window struct {
	*A
	Window *service.Window
}

// NewWindow returns a window accessory.
func NewWindow(info Info) *Window {
	a := Window{}
	a.A = New(info, TypeWindow)

	a.Window = service.NewWindow()
	a.AddS(a.Window.S)

	return &a
}
