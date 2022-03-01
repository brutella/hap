package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeWaterLevel = "B5"

type WaterLevel struct {
	*Float
}

func NewWaterLevel() *WaterLevel {
	c := NewFloat(TypeWaterLevel)
	c.Format = FormatFloat
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(100)

	c.SetValue(0)
	c.Unit = UnitPercentage

	return &WaterLevel{c}
}
