// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeFaucet = "D7"

type Faucet struct {
	*S

	Active *characteristic.Active
}

func NewFaucet() *Faucet {
	s := Faucet{}
	s.S = New(TypeFaucet)

	s.Active = characteristic.NewActive()
	s.AddC(s.Active.C)

	return &s
}
