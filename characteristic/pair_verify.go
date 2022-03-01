package characteristic

// THIS FILE IS AUTO-GENERATED

const TypePairVerify = "4E"

type PairVerify struct {
	*Bytes
}

func NewPairVerify() *PairVerify {
	c := NewBytes(TypePairVerify)
	c.Format = FormatTLV8
	c.Permissions = []string{PermissionRead, PermissionWrite}

	c.SetValue([]byte{})

	return &PairVerify{c}
}
