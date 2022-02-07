// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeLockMechanism = "45"

type LockMechanism struct {
	*S

	LockCurrentState *characteristic.LockCurrentState
	LockTargetState  *characteristic.LockTargetState
}

func NewLockMechanism() *LockMechanism {
	s := LockMechanism{}
	s.S = New(TypeLockMechanism)

	s.LockCurrentState = characteristic.NewLockCurrentState()
	s.AddC(s.LockCurrentState.C)

	s.LockTargetState = characteristic.NewLockTargetState()
	s.AddC(s.LockTargetState.C)

	return &s
}
