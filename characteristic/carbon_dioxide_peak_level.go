package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeCarbonDioxidePeakLevel = "94"

type CarbonDioxidePeakLevel struct {
	*Float
}

func NewCarbonDioxidePeakLevel() *CarbonDioxidePeakLevel {
	c := NewFloat(TypeCarbonDioxidePeakLevel)
	c.Format = FormatFloat
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(100000)

	c.SetValue(0)

	return &CarbonDioxidePeakLevel{c}
}
