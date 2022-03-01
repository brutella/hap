package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeHeatingThresholdTemperature = "12"

type HeatingThresholdTemperature struct {
	*Float
}

func NewHeatingThresholdTemperature() *HeatingThresholdTemperature {
	c := NewFloat(TypeHeatingThresholdTemperature)
	c.Format = FormatFloat
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(25)
	c.SetStepValue(0.1)
	c.SetValue(0)
	c.Unit = UnitCelsius

	return &HeatingThresholdTemperature{c}
}
