package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	CurrentHeatingCoolingStateOff  int = 0
	CurrentHeatingCoolingStateHeat int = 1
	CurrentHeatingCoolingStateCool int = 2
)

const TypeCurrentHeatingCoolingState = "F"

type CurrentHeatingCoolingState struct {
	*Int
}

func NewCurrentHeatingCoolingState() *CurrentHeatingCoolingState {
	c := NewInt(TypeCurrentHeatingCoolingState)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(0)

	return &CurrentHeatingCoolingState{c}
}
