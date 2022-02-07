// THIS FILE IS AUTO-GENERATED
package characteristic

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
