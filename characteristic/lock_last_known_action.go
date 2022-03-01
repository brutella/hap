package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	LockLastKnownActionSecuredPhysicallyInterior   int = 0
	LockLastKnownActionUnsecuredPhysicallyInterior int = 1
	LockLastKnownActionSecuredPhysicallyExterior   int = 2
	LockLastKnownActionUnsecuredPhysicallyExterior int = 3
	LockLastKnownActionSecuredByKeypad             int = 4
	LockLastKnownActionUnsecuredByKeypad           int = 5
	LockLastKnownActionSecuredRemotely             int = 6
	LockLastKnownActionUnsecuredRemotely           int = 7
	LockLastKnownActionSecuredByAutoSecureTimeout  int = 8
)

const TypeLockLastKnownAction = "1C"

type LockLastKnownAction struct {
	*Int
}

func NewLockLastKnownAction() *LockLastKnownAction {
	c := NewInt(TypeLockLastKnownAction)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(0)

	return &LockLastKnownAction{c}
}
