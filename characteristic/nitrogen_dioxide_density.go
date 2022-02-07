// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeNitrogenDioxideDensity = "C4"

type NitrogenDioxideDensity struct {
	*Float
}

func NewNitrogenDioxideDensity() *NitrogenDioxideDensity {
	c := NewFloat(TypeNitrogenDioxideDensity)
	c.Format = FormatFloat
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(1000)
	c.SetStepValue(1)
	c.SetValue(0)

	return &NitrogenDioxideDensity{c}
}
