package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeHue = "13"

type Hue struct {
	*Float
}

func NewHue() *Hue {
	c := NewFloat(TypeHue)
	c.Format = FormatFloat
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(360)
	c.SetStepValue(1)
	c.SetValue(0)
	c.Unit = UnitArcDegrees

	return &Hue{c}
}
