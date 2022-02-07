// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeTelevision = "D8"

type Television struct {
	*S

	Active             *characteristic.Active
	ActiveIdentifier   *characteristic.ActiveIdentifier
	ConfiguredName     *characteristic.ConfiguredName
	SleepDiscoveryMode *characteristic.SleepDiscoveryMode
}

func NewTelevision() *Television {
	s := Television{}
	s.S = New(TypeTelevision)

	s.Active = characteristic.NewActive()
	s.AddC(s.Active.C)

	s.ActiveIdentifier = characteristic.NewActiveIdentifier()
	s.AddC(s.ActiveIdentifier.C)

	s.ConfiguredName = characteristic.NewConfiguredName()
	s.AddC(s.ConfiguredName.C)

	s.SleepDiscoveryMode = characteristic.NewSleepDiscoveryMode()
	s.AddC(s.SleepDiscoveryMode.C)

	return &s
}
