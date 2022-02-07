// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	SlatTypeHorizontal int = 0
	SlatTypeVertical   int = 1
)

const TypeSlatType = "C0"

type SlatType struct {
	*Int
}

func NewSlatType() *SlatType {
	c := NewInt(TypeSlatType)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead}

	c.SetValue(0)

	return &SlatType{c}
}
