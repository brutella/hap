package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeRotationSpeed = "29"

type RotationSpeed struct {
	*Float
}

func NewRotationSpeed() *RotationSpeed {
	c := NewFloat(TypeRotationSpeed)
	c.Format = FormatFloat
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(100)
	c.SetStepValue(1)
	c.SetValue(0)
	c.Unit = UnitPercentage

	return &RotationSpeed{c}
}
