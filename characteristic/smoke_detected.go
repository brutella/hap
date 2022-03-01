package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	SmokeDetectedSmokeNotDetected int = 0
	SmokeDetectedSmokeDetected    int = 1
)

const TypeSmokeDetected = "76"

type SmokeDetected struct {
	*Int
}

func NewSmokeDetected() *SmokeDetected {
	c := NewInt(TypeSmokeDetected)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(0)

	return &SmokeDetected{c}
}
