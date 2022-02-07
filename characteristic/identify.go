// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeIdentify = "14"

type Identify struct {
	*Bool
}

func NewIdentify() *Identify {
	c := NewBool(TypeIdentify)
	c.Format = FormatBool
	c.Permissions = []string{PermissionWrite}

	return &Identify{c}
}
