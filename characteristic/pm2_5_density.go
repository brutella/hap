package characteristic

// THIS FILE IS AUTO-GENERATED

const TypePM2_5Density = "C6"

type PM2_5Density struct {
	*Float
}

func NewPM2_5Density() *PM2_5Density {
	c := NewFloat(TypePM2_5Density)
	c.Format = FormatFloat
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(1000)
	c.SetStepValue(1)
	c.SetValue(0)

	return &PM2_5Density{c}
}
