package characteristic

// THIS FILE IS AUTO-GENERATED

const TypePairingPairings = "50"

type PairingPairings struct {
	*Bytes
}

func NewPairingPairings() *PairingPairings {
	c := NewBytes(TypePairingPairings)
	c.Format = FormatTLV8
	c.Permissions = []string{PermissionRead, PermissionWrite}

	c.SetValue([]byte{})

	return &PairingPairings{c}
}
