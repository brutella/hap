package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	ActiveInactive int = 0
	ActiveActive   int = 1
)

const TypeActive = "B0"

type Active struct {
	*Int
}

func NewActive() *Active {
	c := NewInt(TypeActive)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}

	c.SetValue(0)

	return &Active{c}
}
