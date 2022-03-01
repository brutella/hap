package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeTimeUpdate = "9A"

type TimeUpdate struct {
	*Bool
}

func NewTimeUpdate() *TimeUpdate {
	c := NewBool(TypeTimeUpdate)
	c.Format = FormatBool
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.Val = false

	return &TimeUpdate{c}
}
