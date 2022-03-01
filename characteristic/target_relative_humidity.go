package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeTargetRelativeHumidity = "34"

type TargetRelativeHumidity struct {
	*Float
}

func NewTargetRelativeHumidity() *TargetRelativeHumidity {
	c := NewFloat(TypeTargetRelativeHumidity)
	c.Format = FormatFloat
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(100)
	c.SetStepValue(1)
	c.SetValue(0)
	c.Unit = UnitPercentage

	return &TargetRelativeHumidity{c}
}
