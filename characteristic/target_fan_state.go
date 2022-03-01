package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	TargetFanStateManual int = 0
	TargetFanStateAuto   int = 1
)

const TypeTargetFanState = "BF"

type TargetFanState struct {
	*Int
}

func NewTargetFanState() *TargetFanState {
	c := NewInt(TypeTargetFanState)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}

	c.SetValue(0)

	return &TargetFanState{c}
}
