package accessory

import (
	"github.com/brutella/hap/service"
)

type Outlet struct {
	*A
	Outlet *service.Outlet
}

// NewOutlet returns an outlet accessory.
func NewOutlet(info Info) *Outlet {
	a := Outlet{}
	a.A = New(info, TypeOutlet)

	a.Outlet = service.NewOutlet()
	a.AddS(a.Outlet.S)

	return &a
}
