package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeRelativeHumidityHumidifierThreshold = "CA"

type RelativeHumidityHumidifierThreshold struct {
	*Float
}

func NewRelativeHumidityHumidifierThreshold() *RelativeHumidityHumidifierThreshold {
	c := NewFloat(TypeRelativeHumidityHumidifierThreshold)
	c.Format = FormatFloat
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(100)
	c.SetStepValue(1)
	c.SetValue(0)
	c.Unit = UnitPercentage

	return &RelativeHumidityHumidifierThreshold{c}
}
