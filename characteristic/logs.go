package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeLogs = "1F"

type Logs struct {
	*Bytes
}

func NewLogs() *Logs {
	c := NewBytes(TypeLogs)
	c.Format = FormatTLV8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue([]byte{})

	return &Logs{c}
}
