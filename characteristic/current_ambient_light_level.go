package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeCurrentAmbientLightLevel = "6B"

type CurrentAmbientLightLevel struct {
	*Float
}

func NewCurrentAmbientLightLevel() *CurrentAmbientLightLevel {
	c := NewFloat(TypeCurrentAmbientLightLevel)
	c.Format = FormatFloat
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(0.0001)
	c.SetMaxValue(100000)

	c.SetValue(0.0001)
	c.Unit = UnitLux

	return &CurrentAmbientLightLevel{c}
}
