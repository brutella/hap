package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeTunneledAccessoryAdvertising = "60"

type TunneledAccessoryAdvertising struct {
	*Bool
}

func NewTunneledAccessoryAdvertising() *TunneledAccessoryAdvertising {
	c := NewBool(TypeTunneledAccessoryAdvertising)
	c.Format = FormatBool
	c.Permissions = []string{PermissionWrite, PermissionRead, PermissionEvents}
	c.Val = false

	return &TunneledAccessoryAdvertising{c}
}
