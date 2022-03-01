package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeSaturation = "2F"

type Saturation struct {
	*Float
}

func NewSaturation() *Saturation {
	c := NewFloat(TypeSaturation)
	c.Format = FormatFloat
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(100)
	c.SetStepValue(1)
	c.SetValue(0)
	c.Unit = UnitPercentage

	return &Saturation{c}
}
