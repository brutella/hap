package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeCoolingThresholdTemperature = "D"

type CoolingThresholdTemperature struct {
	*Float
}

func NewCoolingThresholdTemperature() *CoolingThresholdTemperature {
	c := NewFloat(TypeCoolingThresholdTemperature)
	c.Format = FormatFloat
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}
	c.SetMinValue(10)
	c.SetMaxValue(35)
	c.SetStepValue(0.1)
	c.SetValue(10)
	c.Unit = UnitCelsius

	return &CoolingThresholdTemperature{c}
}
