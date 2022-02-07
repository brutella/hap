// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeCurrentTiltAngle = "C1"

type CurrentTiltAngle struct {
	*Int
}

func NewCurrentTiltAngle() *CurrentTiltAngle {
	c := NewInt(TypeCurrentTiltAngle)
	c.Format = FormatInt32
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(-90)
	c.SetMaxValue(90)
	c.SetStepValue(1)
	c.SetValue(-90)
	c.Unit = UnitArcDegrees

	return &CurrentTiltAngle{c}
}
