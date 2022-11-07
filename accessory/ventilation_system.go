package accessory

import (
	"github.com/brutella/hap/service"
)

type VentilationSystem struct {
	*A
	VentilationSystem *service.VentilationSystem
}

// NewVentilationSystem returns an ventilation system accessory.
func NewVentilationSystem(info Info) *VentilationSystem {
	a := VentilationSystem{}
	a.A = New(info, TypeAirPurifier)

	a.VentilationSystem = service.NewVentilationSystem()
	a.AddS(a.VentilationSystem.S)

	return &a
}
