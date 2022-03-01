package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeRelativeHumidityDehumidifierThreshold = "C9"

type RelativeHumidityDehumidifierThreshold struct {
	*Float
}

func NewRelativeHumidityDehumidifierThreshold() *RelativeHumidityDehumidifierThreshold {
	c := NewFloat(TypeRelativeHumidityDehumidifierThreshold)
	c.Format = FormatFloat
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(100)
	c.SetStepValue(1)
	c.SetValue(0)
	c.Unit = UnitPercentage

	return &RelativeHumidityDehumidifierThreshold{c}
}
