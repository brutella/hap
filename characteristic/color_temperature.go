// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeColorTemperature = "CE"

type ColorTemperature struct {
	*Int
}

func NewColorTemperature() *ColorTemperature {
	c := NewInt(TypeColorTemperature)
	c.Format = FormatUInt32
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}
	c.SetMinValue(140)
	c.SetMaxValue(500)
	c.SetStepValue(1)
	c.SetValue(140)

	return &ColorTemperature{c}
}
