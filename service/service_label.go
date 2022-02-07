// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeServiceLabel = "CC"

type ServiceLabel struct {
	*S

	ServiceLabelNamespace *characteristic.ServiceLabelNamespace
}

func NewServiceLabel() *ServiceLabel {
	s := ServiceLabel{}
	s.S = New(TypeServiceLabel)

	s.ServiceLabelNamespace = characteristic.NewServiceLabelNamespace()
	s.AddC(s.ServiceLabelNamespace.C)

	return &s
}
