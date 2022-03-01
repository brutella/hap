package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeBatteryLevel = "68"

type BatteryLevel struct {
	*Int
}

func NewBatteryLevel() *BatteryLevel {
	c := NewInt(TypeBatteryLevel)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(100)
	c.SetStepValue(1)
	c.SetValue(0)
	c.Unit = UnitPercentage

	return &BatteryLevel{c}
}
