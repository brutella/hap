// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeModel = "21"

type Model struct {
	*String
}

func NewModel() *Model {
	c := NewString(TypeModel)
	c.Format = FormatString
	c.Permissions = []string{PermissionRead}

	c.SetValue("")

	return &Model{c}
}
