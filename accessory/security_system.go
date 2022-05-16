package accessory

import (
	"github.com/brutella/hap/service"
)

type SecuritySystem struct {
	*A
	SecuritySystem *service.SecuritySystem
}

// NewSecuritySystem returns a security system accessory.
func NewSecuritySystem(info Info) *SecuritySystem {
	a := SecuritySystem{}
	a.A = New(info, TypeSecuritySystem)

	a.SecuritySystem = service.NewSecuritySystem()
	a.AddS(a.SecuritySystem.S)

	return &a
}
