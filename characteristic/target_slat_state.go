package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	TargetSlatStateManual int = 0
	TargetSlatStateAuto   int = 1
)

const TypeTargetSlatState = "BE"

type TargetSlatState struct {
	*Int
}

func NewTargetSlatState() *TargetSlatState {
	c := NewInt(TypeTargetSlatState)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}

	c.SetValue(0)

	return &TargetSlatState{c}
}
