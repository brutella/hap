package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	LockPhysicalControlsControlLockDisabled int = 0
	LockPhysicalControlsControlLockEnabled  int = 1
)

const TypeLockPhysicalControls = "A7"

type LockPhysicalControls struct {
	*Int
}

func NewLockPhysicalControls() *LockPhysicalControls {
	c := NewInt(TypeLockPhysicalControls)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}

	c.SetValue(0)

	return &LockPhysicalControls{c}
}
