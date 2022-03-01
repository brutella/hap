package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	LockTargetStateUnsecured int = 0
	LockTargetStateSecured   int = 1
)

const TypeLockTargetState = "1E"

type LockTargetState struct {
	*Int
}

func NewLockTargetState() *LockTargetState {
	c := NewInt(TypeLockTargetState)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}

	c.SetValue(0)

	return &LockTargetState{c}
}
