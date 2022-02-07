// THIS FILE IS AUTO-GENERATED
package characteristic

const TypePairSetup = "4C"

type PairSetup struct {
	*Bytes
}

func NewPairSetup() *PairSetup {
	c := NewBytes(TypePairSetup)
	c.Format = FormatTLV8
	c.Permissions = []string{PermissionRead, PermissionWrite}

	c.SetValue([]byte{})

	return &PairSetup{c}
}
