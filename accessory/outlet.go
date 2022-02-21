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
	a := New(info, TypeOutlet)

	outlet := service.NewOutlet()
	a.Ss = append(a.Ss, outlet.S)

	return &Outlet{
		A:      a,
		Outlet: outlet,
	}
}
