package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeTunnelConnectionTimeout = "61"

type TunnelConnectionTimeout struct {
	*Int
}

func NewTunnelConnectionTimeout() *TunnelConnectionTimeout {
	c := NewInt(TypeTunnelConnectionTimeout)
	c.Format = FormatUInt32
	c.Permissions = []string{PermissionWrite, PermissionRead, PermissionEvents}
	c.Val = 0

	return &TunnelConnectionTimeout{c}
}
