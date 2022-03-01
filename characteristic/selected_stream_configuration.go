package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeSelectedStreamConfiguration = "117"

type SelectedStreamConfiguration struct {
	*Bytes
}

func NewSelectedStreamConfiguration() *SelectedStreamConfiguration {
	c := NewBytes(TypeSelectedStreamConfiguration)
	c.Format = FormatTLV8
	c.Permissions = []string{PermissionRead, PermissionWrite}
	c.Val = []byte{}

	return &SelectedStreamConfiguration{c}
}
