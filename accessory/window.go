package accessory

import (
	"github.com/brutella/hap/service"
)

type Windows struct {
	*A
	Window *service.Window
}

// NewWindow returns a window accessory.
func NewWindow(info Info) *Windows {
	a := Windows{}
	a.A = New(info, TypeWindow)
	a.Window = service.NewWindow()
	a.AddS(a.Window.S)

	return &a
}
