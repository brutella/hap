package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeTargetTemperature = "35"

type TargetTemperature struct {
	*Float
}

func NewTargetTemperature() *TargetTemperature {
	c := NewFloat(TypeTargetTemperature)
	c.Format = FormatFloat
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}
	c.SetMinValue(10)
	c.SetMaxValue(38)
	c.SetStepValue(0.1)
	c.SetValue(10)
	c.Unit = UnitCelsius

	return &TargetTemperature{c}
}
