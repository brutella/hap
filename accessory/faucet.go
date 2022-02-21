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
	a := New(info, TypeFaucet)

	faucet := service.NewFaucet()
	a.Ss = append(a.Ss, faucet.S)

	return &Faucet{
		A:      a,
		Faucet: faucet,
	}
}
