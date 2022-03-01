package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeVersion = "37"

type Version struct {
	*String
}

func NewVersion() *Version {
	c := NewString(TypeVersion)
	c.Format = FormatString
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue("")

	return &Version{c}
}
