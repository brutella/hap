package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeProtocolInformation = "A2"

type ProtocolInformation struct {
	*S

	Version *characteristic.Version
}

func NewProtocolInformation() *ProtocolInformation {
	s := ProtocolInformation{}
	s.S = New(TypeProtocolInformation)

	s.Version = characteristic.NewVersion()
	s.AddC(s.Version.C)

	return &s
}
