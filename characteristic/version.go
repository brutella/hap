// THIS FILE IS AUTO-GENERATED
package characteristic

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
