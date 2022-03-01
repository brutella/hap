package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	TargetHumidifierDehumidifierStateHumidifierOrDehumidifier int = 0
	TargetHumidifierDehumidifierStateHumidifier               int = 1
	TargetHumidifierDehumidifierStateDehumidifier             int = 2
)

const TypeTargetHumidifierDehumidifierState = "B4"

type TargetHumidifierDehumidifierState struct {
	*Int
}

func NewTargetHumidifierDehumidifierState() *TargetHumidifierDehumidifierState {
	c := NewInt(TypeTargetHumidifierDehumidifierState)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}

	c.SetValue(0)

	return &TargetHumidifierDehumidifierState{c}
}
