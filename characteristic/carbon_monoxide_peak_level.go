package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeCarbonMonoxidePeakLevel = "91"

type CarbonMonoxidePeakLevel struct {
	*Float
}

func NewCarbonMonoxidePeakLevel() *CarbonMonoxidePeakLevel {
	c := NewFloat(TypeCarbonMonoxidePeakLevel)
	c.Format = FormatFloat
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(100)

	c.SetValue(0)

	return &CarbonMonoxidePeakLevel{c}
}
