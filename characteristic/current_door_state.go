package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	CurrentDoorStateOpen    int = 0
	CurrentDoorStateClosed  int = 1
	CurrentDoorStateOpening int = 2
	CurrentDoorStateClosing int = 3
	CurrentDoorStateStopped int = 4
)

const TypeCurrentDoorState = "E"

type CurrentDoorState struct {
	*Int
}

func NewCurrentDoorState() *CurrentDoorState {
	c := NewInt(TypeCurrentDoorState)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(0)

	return &CurrentDoorState{c}
}
