package accessory

import (
	"github.com/brutella/hap/service"
)

type Television struct {
	*A
	Television *service.Television
	Speaker    *service.Speaker
}

// NewTelevision returns a television accessory.
func NewTelevision(info Info) *Television {
	a := Television{}
	a.A = New(info, TypeTelevision)

	a.Television = service.NewTelevision()
	a.AddS(a.Television.S)

	a.Speaker = service.NewSpeaker()
	a.AddS(a.Speaker.S)

	return &a
}
