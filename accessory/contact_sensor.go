package accessory

import (
	"github.com/brutella/hap/service"
)

type ContactSensor struct {
	*A
	ContactSensor *service.ContactSensor
}

// NewContactSensor returns a ContactSensor which implements model.ContactSensor.
func NewContactSensor(info Info) *ContactSensor {
	a := ContactSensor{}
	a.A = New(info, TypeSensor)

	a.ContactSensor = service.NewContactSensor()
	a.AddS(a.ContactSensor.S)

	return &a
}
