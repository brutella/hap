package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeBrightness = "8"

type Brightness struct {
	*Int
}

func NewBrightness() *Brightness {
	c := NewInt(TypeBrightness)
	c.Format = FormatInt32
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(100)
	c.SetStepValue(1)
	c.SetValue(0)
	c.Unit = UnitPercentage

	return &Brightness{c}
}
