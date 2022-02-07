// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeCurrentRelativeHumidity = "10"

type CurrentRelativeHumidity struct {
	*Float
}

func NewCurrentRelativeHumidity() *CurrentRelativeHumidity {
	c := NewFloat(TypeCurrentRelativeHumidity)
	c.Format = FormatFloat
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(100)
	c.SetStepValue(1)
	c.SetValue(0)
	c.Unit = UnitPercentage

	return &CurrentRelativeHumidity{c}
}
