package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	TargetHeaterCoolerStateAuto int = 0
	TargetHeaterCoolerStateHeat int = 1
	TargetHeaterCoolerStateCool int = 2
)

const TypeTargetHeaterCoolerState = "B2"

type TargetHeaterCoolerState struct {
	*Int
}

func NewTargetHeaterCoolerState() *TargetHeaterCoolerState {
	c := NewInt(TypeTargetHeaterCoolerState)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}

	c.SetValue(0)

	return &TargetHeaterCoolerState{c}
}
