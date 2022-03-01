package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	CarbonMonoxideDetectedCOLevelsNormal   int = 0
	CarbonMonoxideDetectedCOLevelsAbnormal int = 1
)

const TypeCarbonMonoxideDetected = "69"

type CarbonMonoxideDetected struct {
	*Int
}

func NewCarbonMonoxideDetected() *CarbonMonoxideDetected {
	c := NewInt(TypeCarbonMonoxideDetected)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(0)

	return &CarbonMonoxideDetected{c}
}
