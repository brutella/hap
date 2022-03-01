package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	ChargingStateNotCharging   int = 0
	ChargingStateCharging      int = 1
	ChargingStateNotChargeable int = 2
)

const TypeChargingState = "8F"

type ChargingState struct {
	*Int
}

func NewChargingState() *ChargingState {
	c := NewInt(TypeChargingState)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(0)

	return &ChargingState{c}
}
