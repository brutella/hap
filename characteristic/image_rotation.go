package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeImageRotation = "11E"

type ImageRotation struct {
	*Float
}

func NewImageRotation() *ImageRotation {
	c := NewFloat(TypeImageRotation)
	c.Format = FormatFloat
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(270)
	c.SetStepValue(90)
	c.SetValue(0)
	c.Unit = UnitArcDegrees

	return &ImageRotation{c}
}
