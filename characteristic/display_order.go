package characteristic

// THIS FILE IS AUTO-GENERATED

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
