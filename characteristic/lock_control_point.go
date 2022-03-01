package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeLockControlPoint = "19"

type LockControlPoint struct {
	*Bytes
}

func NewLockControlPoint() *LockControlPoint {
	c := NewBytes(TypeLockControlPoint)
	c.Format = FormatTLV8
	c.Permissions = []string{PermissionWrite}

	return &LockControlPoint{c}
}
