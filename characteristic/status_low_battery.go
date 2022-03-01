package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	StatusLowBatteryBatteryLevelNormal int = 0
	StatusLowBatteryBatteryLevelLow    int = 1
)

const TypeStatusLowBattery = "79"

type StatusLowBattery struct {
	*Int
}

func NewStatusLowBattery() *StatusLowBattery {
	c := NewInt(TypeStatusLowBattery)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(0)

	return &StatusLowBattery{c}
}
