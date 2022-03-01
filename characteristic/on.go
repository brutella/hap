package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeOn = "25"

type On struct {
	*Bool
}

func NewOn() *On {
	c := NewBool(TypeOn)
	c.Format = FormatBool
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}

	c.SetValue(false)

	return &On{c}
}
