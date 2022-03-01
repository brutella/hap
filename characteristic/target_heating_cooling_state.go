package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	TargetHeatingCoolingStateOff  int = 0
	TargetHeatingCoolingStateHeat int = 1
	TargetHeatingCoolingStateCool int = 2
	TargetHeatingCoolingStateAuto int = 3
)

const TypeTargetHeatingCoolingState = "33"

type TargetHeatingCoolingState struct {
	*Int
}

func NewTargetHeatingCoolingState() *TargetHeatingCoolingState {
	c := NewInt(TypeTargetHeatingCoolingState)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}

	c.SetValue(0)

	return &TargetHeatingCoolingState{c}
}
