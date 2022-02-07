// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeDisplayOrder = "136"

type DisplayOrder struct {
	*Bytes
}

func NewDisplayOrder() *DisplayOrder {
	c := NewBytes(TypeDisplayOrder)
	c.Format = FormatTLV8
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}

	c.SetValue([]byte{})

	return &DisplayOrder{c}
}
