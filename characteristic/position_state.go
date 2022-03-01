package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	PositionStateDecreasing int = 0
	PositionStateIncreasing int = 1
	PositionStateStopped    int = 2
)

const TypePositionState = "72"

type PositionState struct {
	*Int
}

func NewPositionState() *PositionState {
	c := NewInt(TypePositionState)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(0)

	return &PositionState{c}
}
