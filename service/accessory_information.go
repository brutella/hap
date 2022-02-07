// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeAccessoryInformation = "3E"

type AccessoryInformation struct {
	*S

	Identify         *characteristic.Identify
	Manufacturer     *characteristic.Manufacturer
	Model            *characteristic.Model
	Name             *characteristic.Name
	SerialNumber     *characteristic.SerialNumber
	FirmwareRevision *characteristic.FirmwareRevision
}

func NewAccessoryInformation() *AccessoryInformation {
	s := AccessoryInformation{}
	s.S = New(TypeAccessoryInformation)

	s.Identify = characteristic.NewIdentify()
	s.AddC(s.Identify.C)

	s.Manufacturer = characteristic.NewManufacturer()
	s.AddC(s.Manufacturer.C)

	s.Model = characteristic.NewModel()
	s.AddC(s.Model.C)

	s.Name = characteristic.NewName()
	s.AddC(s.Name.C)

	s.SerialNumber = characteristic.NewSerialNumber()
	s.AddC(s.SerialNumber.C)

	s.FirmwareRevision = characteristic.NewFirmwareRevision()
	s.AddC(s.FirmwareRevision.C)

	return &s
}
