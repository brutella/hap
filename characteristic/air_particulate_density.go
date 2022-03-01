package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeAirParticulateDensity = "64"

type AirParticulateDensity struct {
	*Float
}

func NewAirParticulateDensity() *AirParticulateDensity {
	c := NewFloat(TypeAirParticulateDensity)
	c.Format = FormatFloat
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(1000)
	c.SetStepValue(1)
	c.SetValue(0)

	return &AirParticulateDensity{c}
}
