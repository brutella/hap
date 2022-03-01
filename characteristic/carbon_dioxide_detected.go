package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	CarbonDioxideDetectedCO2LevelsNormal   int = 0
	CarbonDioxideDetectedCO2LevelsAbnormal int = 1
)

const TypeCarbonDioxideDetected = "92"

type CarbonDioxideDetected struct {
	*Int
}

func NewCarbonDioxideDetected() *CarbonDioxideDetected {
	c := NewInt(TypeCarbonDioxideDetected)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(0)

	return &CarbonDioxideDetected{c}
}
