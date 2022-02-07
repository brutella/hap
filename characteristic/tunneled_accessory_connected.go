// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeTunneledAccessoryConnected = "59"

type TunneledAccessoryConnected struct {
	*Bool
}

func NewTunneledAccessoryConnected() *TunneledAccessoryConnected {
	c := NewBool(TypeTunneledAccessoryConnected)
	c.Format = FormatBool
	c.Permissions = []string{PermissionWrite, PermissionRead, PermissionEvents}
	c.Val = false

	return &TunneledAccessoryConnected{c}
}
