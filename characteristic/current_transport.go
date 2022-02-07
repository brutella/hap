package characteristic

const TypeCurrentTransport = "22B"

type CurrentTransport struct {
	*Bool
}

func NewCurrentTransport() *CurrentTransport {
	c := NewBool(TypeCurrentTransport)
	c.Format = FormatBool
	c.Permissions = []string{PermissionRead}
	c.Val = false

	return &CurrentTransport{c}
}
