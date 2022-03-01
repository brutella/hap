package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeLockManagementAutoSecurityTimeout = "1A"

type LockManagementAutoSecurityTimeout struct {
	*Int
}

func NewLockManagementAutoSecurityTimeout() *LockManagementAutoSecurityTimeout {
	c := NewInt(TypeLockManagementAutoSecurityTimeout)
	c.Format = FormatUInt32
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}

	c.SetValue(0)
	c.Unit = UnitSeconds

	return &LockManagementAutoSecurityTimeout{c}
}
