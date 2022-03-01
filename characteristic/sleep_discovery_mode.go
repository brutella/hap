package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	SleepDiscoveryModeNotDiscoverable    int = 0
	SleepDiscoveryModeAlwaysDiscoverable int = 1
)

const TypeSleepDiscoveryMode = "E8"

type SleepDiscoveryMode struct {
	*Int
}

func NewSleepDiscoveryMode() *SleepDiscoveryMode {
	c := NewInt(TypeSleepDiscoveryMode)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(1)

	c.SetValue(0)

	return &SleepDiscoveryMode{c}
}
