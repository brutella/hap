// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeName = "23"

type Name struct {
	*String
}

func NewName() *Name {
	c := NewString(TypeName)
	c.Format = FormatString
	c.Permissions = []string{PermissionRead}

	c.SetValue("")

	return &Name{c}
}
