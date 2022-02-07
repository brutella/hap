package service

import (
	"github.com/brutella/hap/characteristic"

	"encoding/json"
)

type S struct {
	Id      uint64
	Type    string
	Hidden  bool
	Primary bool

	Linked []*S
	Cs     []*characteristic.C
}

// New returns a new service.
func New(typ string) *S {
	return &S{
		Type:   typ,
		Cs:     []*characteristic.C{},
		Linked: []*S{},
	}
}

func (s *S) AddC(c *characteristic.C) {
	s.Cs = append(s.Cs, c)
}

func (s *S) AddS(other *S) {
	s.Linked = append(s.Linked, other)
}

func (s *S) C(typ string) *characteristic.C {
	for _, c := range s.Cs {
		if c.Type == typ {
			return c
		}
	}

	return nil
}

func (s *S) MarshalJSON() ([]byte, error) {
	linked := []uint64{}
	for _, s := range s.Linked {
		linked = append(linked, s.Id)
	}

	d := struct {
		Id      uint64              `json:"iid"`
		Type    string              `json:"type"`
		Cs      []*characteristic.C `json:"characteristics"`
		Hidden  *bool               `json:"hidden,omitempty"`
		Primary *bool               `json:"primary,omitempty"`
		Linked  []uint64            `json:"linked,omitempty"`
	}{
		Id:     s.Id,
		Type:   s.Type,
		Cs:     s.Cs,
		Linked: linked,
	}

	if s.Hidden {
		d.Hidden = &s.Hidden
	}

	if s.Primary {
		d.Primary = &s.Primary
	}

	return json.Marshal(&d)
}
