package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeCurrentTemperature = "11"

type CurrentTemperature struct {
	*Float
}

func NewCurrentTemperature() *CurrentTemperature {
	c := NewFloat(TypeCurrentTemperature)
	c.Format = FormatFloat
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(100)
	c.SetStepValue(0.1)
	c.SetValue(0)
	c.Unit = UnitCelsius

	return &CurrentTemperature{c}
}
