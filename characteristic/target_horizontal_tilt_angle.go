package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeTargetHorizontalTiltAngle = "7B"

type TargetHorizontalTiltAngle struct {
	*Int
}

func NewTargetHorizontalTiltAngle() *TargetHorizontalTiltAngle {
	c := NewInt(TypeTargetHorizontalTiltAngle)
	c.Format = FormatInt32
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}
	c.SetMinValue(-90)
	c.SetMaxValue(90)
	c.SetStepValue(1)
	c.SetValue(-90)
	c.Unit = UnitArcDegrees

	return &TargetHorizontalTiltAngle{c}
}
