// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeServiceLabelIndex = "CB"

type ServiceLabelIndex struct {
	*Int
}

func NewServiceLabelIndex() *ServiceLabelIndex {
	c := NewInt(TypeServiceLabelIndex)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead}
	c.SetMinValue(1)
	c.SetMaxValue(255)
	c.SetStepValue(1)
	c.SetValue(1)

	return &ServiceLabelIndex{c}
}
