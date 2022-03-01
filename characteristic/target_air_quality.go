package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	TargetAirQualityExcellent int = 0
	TargetAirQualityGood      int = 1
	TargetAirQualityFair      int = 2
)

const TypeTargetAirQuality = "AE"

type TargetAirQuality struct {
	*Int
}

func NewTargetAirQuality() *TargetAirQuality {
	c := NewInt(TypeTargetAirQuality)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}

	c.SetValue(0)

	return &TargetAirQuality{c}
}
