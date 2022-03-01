package characteristic

// THIS FILE IS AUTO-GENERATED

const TypePM10Density = "C7"

type PM10Density struct {
	*Float
}

func NewPM10Density() *PM10Density {
	c := NewFloat(TypePM10Density)
	c.Format = FormatFloat
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(1000)
	c.SetStepValue(1)
	c.SetValue(0)

	return &PM10Density{c}
}
