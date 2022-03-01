package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	TargetDoorStateOpen   int = 0
	TargetDoorStateClosed int = 1
)

const TypeTargetDoorState = "32"

type TargetDoorState struct {
	*Int
}

func NewTargetDoorState() *TargetDoorState {
	c := NewInt(TypeTargetDoorState)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}

	c.SetValue(0)

	return &TargetDoorState{c}
}
