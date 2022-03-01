package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	LockCurrentStateUnsecured int = 0
	LockCurrentStateSecured   int = 1
	LockCurrentStateJammed    int = 2
	LockCurrentStateUnknown   int = 3
)

const TypeLockCurrentState = "1D"

type LockCurrentState struct {
	*Int
}

func NewLockCurrentState() *LockCurrentState {
	c := NewInt(TypeLockCurrentState)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(0)

	return &LockCurrentState{c}
}
