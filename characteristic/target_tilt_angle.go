package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeTargetTiltAngle = "C2"

type TargetTiltAngle struct {
	*Int
}

func NewTargetTiltAngle() *TargetTiltAngle {
	c := NewInt(TypeTargetTiltAngle)
	c.Format = FormatInt32
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}
	c.SetMinValue(-90)
	c.SetMaxValue(90)
	c.SetStepValue(1)
	c.SetValue(-90)
	c.Unit = UnitArcDegrees

	return &TargetTiltAngle{c}
}
