// THIS FILE IS AUTO-GENERATED
package characteristic

const TypePairingFeatures = "4F"

type PairingFeatures struct {
	*Int
}

func NewPairingFeatures() *PairingFeatures {
	c := NewInt(TypePairingFeatures)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead}

	c.SetValue(0)

	return &PairingFeatures{c}
}
