package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeOzoneDensity = "C3"

type OzoneDensity struct {
	*Float
}

func NewOzoneDensity() *OzoneDensity {
	c := NewFloat(TypeOzoneDensity)
	c.Format = FormatFloat
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(1000)
	c.SetStepValue(1)
	c.SetValue(0)

	return &OzoneDensity{c}
}
