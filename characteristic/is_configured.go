package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	IsConfiguredNotConfigured int = 0
	IsConfiguredConfigured    int = 1
)

const TypeIsConfigured = "D6"

type IsConfigured struct {
	*Int
}

func NewIsConfigured() *IsConfigured {
	c := NewInt(TypeIsConfigured)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}

	c.SetValue(0)

	return &IsConfigured{c}
}
