package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeCarbonMonoxideLevel = "90"

type CarbonMonoxideLevel struct {
	*Float
}

func NewCarbonMonoxideLevel() *CarbonMonoxideLevel {
	c := NewFloat(TypeCarbonMonoxideLevel)
	c.Format = FormatFloat
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(100)

	c.SetValue(0)

	return &CarbonMonoxideLevel{c}
}
