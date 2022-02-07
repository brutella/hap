// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeAppMatchingIdentifier = "A4"

type AppMatchingIdentifier struct {
	*Bytes
}

func NewAppMatchingIdentifier() *AppMatchingIdentifier {
	c := NewBytes(TypeAppMatchingIdentifier)
	c.Format = FormatTLV8
	c.Permissions = []string{PermissionRead}
	c.Val = []byte{}

	return &AppMatchingIdentifier{c}
}
