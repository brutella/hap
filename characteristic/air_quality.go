package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	AirQualityUnknown   int = 0
	AirQualityExcellent int = 1
	AirQualityGood      int = 2
	AirQualityFair      int = 3
	AirQualityInferior  int = 4
	AirQualityPoor      int = 5
)

const TypeAirQuality = "95"

type AirQuality struct {
	*Int
}

func NewAirQuality() *AirQuality {
	c := NewInt(TypeAirQuality)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(0)

	return &AirQuality{c}
}
