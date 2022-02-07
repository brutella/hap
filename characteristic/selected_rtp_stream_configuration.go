// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeSelectedRTPStreamConfiguration = "117"

type SelectedRTPStreamConfiguration struct {
	*Bytes
}

func NewSelectedRTPStreamConfiguration() *SelectedRTPStreamConfiguration {
	c := NewBytes(TypeSelectedRTPStreamConfiguration)
	c.Format = FormatTLV8
	c.Permissions = []string{PermissionRead, PermissionWrite}

	c.SetValue([]byte{})

	return &SelectedRTPStreamConfiguration{c}
}
