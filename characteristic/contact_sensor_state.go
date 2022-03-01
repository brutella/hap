package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	ContactSensorStateContactDetected    int = 0
	ContactSensorStateContactNotDetected int = 1
)

const TypeContactSensorState = "6A"

type ContactSensorState struct {
	*Int
}

func NewContactSensorState() *ContactSensorState {
	c := NewInt(TypeContactSensorState)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(0)

	return &ContactSensorState{c}
}
