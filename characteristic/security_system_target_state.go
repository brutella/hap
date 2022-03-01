package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	SecuritySystemTargetStateStayArm  int = 0
	SecuritySystemTargetStateAwayArm  int = 1
	SecuritySystemTargetStateNightArm int = 2
	SecuritySystemTargetStateDisarm   int = 3
)

const TypeSecuritySystemTargetState = "67"

type SecuritySystemTargetState struct {
	*Int
}

func NewSecuritySystemTargetState() *SecuritySystemTargetState {
	c := NewInt(TypeSecuritySystemTargetState)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}

	c.SetValue(0)

	return &SecuritySystemTargetState{c}
}
