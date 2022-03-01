package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	CurrentFanStateInactive   int = 0
	CurrentFanStateIdle       int = 1
	CurrentFanStateBlowingAir int = 2
)

const TypeCurrentFanState = "AF"

type CurrentFanState struct {
	*Int
}

func NewCurrentFanState() *CurrentFanState {
	c := NewInt(TypeCurrentFanState)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(0)

	return &CurrentFanState{c}
}
