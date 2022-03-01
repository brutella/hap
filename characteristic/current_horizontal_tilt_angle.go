package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeCurrentHorizontalTiltAngle = "6C"

type CurrentHorizontalTiltAngle struct {
	*Int
}

func NewCurrentHorizontalTiltAngle() *CurrentHorizontalTiltAngle {
	c := NewInt(TypeCurrentHorizontalTiltAngle)
	c.Format = FormatInt32
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(-90)
	c.SetMaxValue(90)
	c.SetStepValue(1)
	c.SetValue(-90)
	c.Unit = UnitArcDegrees

	return &CurrentHorizontalTiltAngle{c}
}
