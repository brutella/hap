package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeStatusActive = "75"

type StatusActive struct {
	*Bool
}

func NewStatusActive() *StatusActive {
	c := NewBool(TypeStatusActive)
	c.Format = FormatBool
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(false)

	return &StatusActive{c}
}
