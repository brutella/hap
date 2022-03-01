package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	StatusTamperedNotTampered int = 0
	StatusTamperedTampered    int = 1
)

const TypeStatusTampered = "7A"

type StatusTampered struct {
	*Int
}

func NewStatusTampered() *StatusTampered {
	c := NewInt(TypeStatusTampered)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(0)

	return &StatusTampered{c}
}
