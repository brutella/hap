package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	TargetAirPurifierStateManual int = 0
	TargetAirPurifierStateAuto   int = 1
)

const TypeTargetAirPurifierState = "A8"

type TargetAirPurifierState struct {
	*Int
}

func NewTargetAirPurifierState() *TargetAirPurifierState {
	c := NewInt(TypeTargetAirPurifierState)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}

	c.SetValue(0)

	return &TargetAirPurifierState{c}
}
