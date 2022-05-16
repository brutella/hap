package accessory

import (
	"github.com/brutella/hap/service"
)

type Faucet struct {
	*A
	Faucet *service.Faucet
}

// NewFaucet returns an outlet accessory.
func NewFaucet(info Info) *Faucet {
	a := Faucet{}
	a.A = New(info, TypeFaucet)

	a.Faucet = service.NewFaucet()
	a.AddS(a.Faucet.S)

	return &a
}
