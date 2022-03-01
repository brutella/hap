package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeTunneledAccessoryStateNumber = "58"

type TunneledAccessoryStateNumber struct {
	*Float
}

func NewTunneledAccessoryStateNumber() *TunneledAccessoryStateNumber {
	c := NewFloat(TypeTunneledAccessoryStateNumber)
	c.Format = FormatFloat
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.Val = 0

	return &TunneledAccessoryStateNumber{c}
}
