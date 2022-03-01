package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	StatusJammedNotJammed int = 0
	StatusJammedJammed    int = 1
)

const TypeStatusJammed = "78"

type StatusJammed struct {
	*Int
}

func NewStatusJammed() *StatusJammed {
	c := NewInt(TypeStatusJammed)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(0)

	return &StatusJammed{c}
}
