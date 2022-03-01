package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeVOCDensity = "C8"

type VOCDensity struct {
	*Float
}

func NewVOCDensity() *VOCDensity {
	c := NewFloat(TypeVOCDensity)
	c.Format = FormatFloat
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(1000)
	c.SetStepValue(1)
	c.SetValue(0)

	return &VOCDensity{c}
}
