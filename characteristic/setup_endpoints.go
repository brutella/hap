// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeSetupEndpoints = "118"

type SetupEndpoints struct {
	*Bytes
}

func NewSetupEndpoints() *SetupEndpoints {
	c := NewBytes(TypeSetupEndpoints)
	c.Format = FormatTLV8
	c.Permissions = []string{PermissionRead, PermissionWrite}

	c.SetValue([]byte{})

	return &SetupEndpoints{c}
}
