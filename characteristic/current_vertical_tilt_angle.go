package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeCurrentVerticalTiltAngle = "6E"

type CurrentVerticalTiltAngle struct {
	*Int
}

func NewCurrentVerticalTiltAngle() *CurrentVerticalTiltAngle {
	c := NewInt(TypeCurrentVerticalTiltAngle)
	c.Format = FormatInt32
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(-90)
	c.SetMaxValue(90)
	c.SetStepValue(1)
	c.SetValue(-90)
	c.Unit = UnitArcDegrees

	return &CurrentVerticalTiltAngle{c}
}
