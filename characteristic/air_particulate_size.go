package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	AirParticulateSize2_5Μm int = 0
	AirParticulateSize10Μm  int = 1
)

const TypeAirParticulateSize = "65"

type AirParticulateSize struct {
	*Int
}

func NewAirParticulateSize() *AirParticulateSize {
	c := NewInt(TypeAirParticulateSize)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(0)

	return &AirParticulateSize{c}
}
