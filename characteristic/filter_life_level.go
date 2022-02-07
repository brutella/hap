// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeFilterLifeLevel = "AB"

type FilterLifeLevel struct {
	*Float
}

func NewFilterLifeLevel() *FilterLifeLevel {
	c := NewFloat(TypeFilterLifeLevel)
	c.Format = FormatFloat
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(100)

	c.SetValue(0)

	return &FilterLifeLevel{c}
}
