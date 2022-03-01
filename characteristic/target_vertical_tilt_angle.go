package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeTargetVerticalTiltAngle = "7D"

type TargetVerticalTiltAngle struct {
	*Int
}

func NewTargetVerticalTiltAngle() *TargetVerticalTiltAngle {
	c := NewInt(TypeTargetVerticalTiltAngle)
	c.Format = FormatInt32
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}
	c.SetMinValue(-90)
	c.SetMaxValue(90)
	c.SetStepValue(1)
	c.SetValue(-90)
	c.Unit = UnitArcDegrees

	return &TargetVerticalTiltAngle{c}
}
