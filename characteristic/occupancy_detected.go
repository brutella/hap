package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	OccupancyDetectedOccupancyNotDetected int = 0
	OccupancyDetectedOccupancyDetected    int = 1
)

const TypeOccupancyDetected = "71"

type OccupancyDetected struct {
	*Int
}

func NewOccupancyDetected() *OccupancyDetected {
	c := NewInt(TypeOccupancyDetected)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(0)

	return &OccupancyDetected{c}
}
