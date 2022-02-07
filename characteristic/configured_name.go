// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeConfiguredName = "E3"

type ConfiguredName struct {
	*String
}

func NewConfiguredName() *ConfiguredName {
	c := NewString(TypeConfiguredName)
	c.Format = FormatString
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}

	c.SetValue("")

	return &ConfiguredName{c}
}
