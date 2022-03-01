package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	RotationDirectionClockwise        int = 0
	RotationDirectionCounterclockwise int = 1
)

const TypeRotationDirection = "28"

type RotationDirection struct {
	*Int
}

func NewRotationDirection() *RotationDirection {
	c := NewInt(TypeRotationDirection)
	c.Format = FormatInt32
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}

	c.SetValue(0)

	return &RotationDirection{c}
}
