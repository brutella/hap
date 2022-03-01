package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeSulphurDioxideDensity = "C5"

type SulphurDioxideDensity struct {
	*Float
}

func NewSulphurDioxideDensity() *SulphurDioxideDensity {
	c := NewFloat(TypeSulphurDioxideDensity)
	c.Format = FormatFloat
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(1000)
	c.SetStepValue(1)
	c.SetValue(0)

	return &SulphurDioxideDensity{c}
}
