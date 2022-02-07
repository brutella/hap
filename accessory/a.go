package accessory

import (
	"github.com/brutella/hap/service"

	"encoding/json"
	"net/http"
)

type A struct {
	Id   uint64
	Type byte
	Info *service.AccessoryInformation
	Ss   []*service.S
	// IdentifyFunc is called when a client
	// makes a POST to the /identify endpoint.
	IdentifyFunc func(*http.Request)
}

type Info struct {
	Name         string
	SerialNumber string
	Manufacturer string
	Model        string
	Firmware     string
}

func New(info Info, typ byte) *A {
	s := service.NewAccessoryInformation()
	s.Name.SetValue("-")
	s.Model.SetValue("-")
	s.SerialNumber.SetValue("-")
	s.Manufacturer.SetValue("-")
	s.FirmwareRevision.SetValue("-")

	if info.Name != "" {
		s.Name.Val = info.Name
	}

	if info.Model != "" {
		s.Model.Val = info.Model
	}

	if info.SerialNumber != "" {
		s.SerialNumber.Val = info.SerialNumber
	}

	if info.Manufacturer != "" {
		s.Manufacturer.Val = info.Manufacturer
	}

	if info.Firmware != "" {
		s.FirmwareRevision.Val = info.Firmware
	}

	return &A{
		Type: typ,
		Info: s,
		Ss:   []*service.S{s.S},
	}
}

// Adds a service to the accessory and updates the ids of the service and the corresponding characteristics
func (a *A) AddS(s *service.S) {
	a.Ss = append(a.Ss, s)
}

func (a *A) Name() string {
	return a.Info.Name.Value()
}

func (a *A) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Id uint64       `json:"aid"`
		Ss []*service.S `json:"services"`
	}{
		Id: a.Id,
		Ss: a.Ss,
	})
}
