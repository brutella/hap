package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	InUseNotInUse int = 0
	InUseInUse    int = 1
)

const TypeInUse = "D2"

type InUse struct {
	*Int
}

func NewInUse() *InUse {
	c := NewInt(TypeInUse)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(0)

	return &InUse{c}
}
