package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	CurrentAirPurifierStateInactive     int = 0
	CurrentAirPurifierStateIdle         int = 1
	CurrentAirPurifierStatePurifyingAir int = 2
)

const TypeCurrentAirPurifierState = "A9"

type CurrentAirPurifierState struct {
	*Int
}

func NewCurrentAirPurifierState() *CurrentAirPurifierState {
	c := NewInt(TypeCurrentAirPurifierState)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(0)

	return &CurrentAirPurifierState{c}
}
