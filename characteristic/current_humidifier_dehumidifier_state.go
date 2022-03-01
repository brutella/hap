package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	CurrentHumidifierDehumidifierStateInactive      int = 0
	CurrentHumidifierDehumidifierStateIdle          int = 1
	CurrentHumidifierDehumidifierStateHumidifying   int = 2
	CurrentHumidifierDehumidifierStateDehumidifying int = 3
)

const TypeCurrentHumidifierDehumidifierState = "B3"

type CurrentHumidifierDehumidifierState struct {
	*Int
}

func NewCurrentHumidifierDehumidifierState() *CurrentHumidifierDehumidifierState {
	c := NewInt(TypeCurrentHumidifierDehumidifierState)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(0)

	return &CurrentHumidifierDehumidifierState{c}
}
