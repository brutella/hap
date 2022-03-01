package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	SecuritySystemCurrentStateStayArm        int = 0
	SecuritySystemCurrentStateAwayArm        int = 1
	SecuritySystemCurrentStateNightArm       int = 2
	SecuritySystemCurrentStateDisarmed       int = 3
	SecuritySystemCurrentStateAlarmTriggered int = 4
)

const TypeSecuritySystemCurrentState = "66"

type SecuritySystemCurrentState struct {
	*Int
}

func NewSecuritySystemCurrentState() *SecuritySystemCurrentState {
	c := NewInt(TypeSecuritySystemCurrentState)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(0)

	return &SecuritySystemCurrentState{c}
}
