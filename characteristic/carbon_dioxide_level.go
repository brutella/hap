// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeCarbonDioxideLevel = "93"

type CarbonDioxideLevel struct {
	*Float
}

func NewCarbonDioxideLevel() *CarbonDioxideLevel {
	c := NewFloat(TypeCarbonDioxideLevel)
	c.Format = FormatFloat
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(100000)

	c.SetValue(0)

	return &CarbonDioxideLevel{c}
}
