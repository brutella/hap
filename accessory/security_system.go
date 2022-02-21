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
	a := New(info, TypeSecuritySystem)

	garage := service.NewSecuritySystem()
	a.Ss = append(a.Ss, garage.S)

	return &SecuritySystem{
		A:              a,
		SecuritySystem: garage,
	}
}
