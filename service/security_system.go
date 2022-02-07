// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeSecuritySystem = "7E"

type SecuritySystem struct {
	*S

	SecuritySystemCurrentState *characteristic.SecuritySystemCurrentState
	SecuritySystemTargetState  *characteristic.SecuritySystemTargetState
}

func NewSecuritySystem() *SecuritySystem {
	s := SecuritySystem{}
	s.S = New(TypeSecuritySystem)

	s.SecuritySystemCurrentState = characteristic.NewSecuritySystemCurrentState()
	s.AddC(s.SecuritySystemCurrentState.C)

	s.SecuritySystemTargetState = characteristic.NewSecuritySystemTargetState()
	s.AddC(s.SecuritySystemTargetState.C)

	return &s
}
