// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeLockManagement = "44"

type LockManagement struct {
	*S

	LockControlPoint *characteristic.LockControlPoint
	Version          *characteristic.Version
}

func NewLockManagement() *LockManagement {
	s := LockManagement{}
	s.S = New(TypeLockManagement)

	s.LockControlPoint = characteristic.NewLockControlPoint()
	s.AddC(s.LockControlPoint.C)

	s.Version = characteristic.NewVersion()
	s.AddC(s.Version.C)

	return &s
}
