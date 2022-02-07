// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeInputSource = "D9"

type InputSource struct {
	*S

	ConfiguredName         *characteristic.ConfiguredName
	InputSourceType        *characteristic.InputSourceType
	IsConfigured           *characteristic.IsConfigured
	CurrentVisibilityState *characteristic.CurrentVisibilityState
}

func NewInputSource() *InputSource {
	s := InputSource{}
	s.S = New(TypeInputSource)

	s.ConfiguredName = characteristic.NewConfiguredName()
	s.AddC(s.ConfiguredName.C)

	s.InputSourceType = characteristic.NewInputSourceType()
	s.AddC(s.InputSourceType.C)

	s.IsConfigured = characteristic.NewIsConfigured()
	s.AddC(s.IsConfigured.C)

	s.CurrentVisibilityState = characteristic.NewCurrentVisibilityState()
	s.AddC(s.CurrentVisibilityState.C)

	return &s
}
