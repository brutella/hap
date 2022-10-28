package characteristic

const TypePing = "23C"

type Ping struct {
	*Bytes
}

func NewPing() *Ping {
	c := NewBytes(TypePing)
	c.Format = FormatTLV8
	c.Permissions = []string{PermissionRead}

	c.SetValue([]byte{})

	return &Ping{c}
}
