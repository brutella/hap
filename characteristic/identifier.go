package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeIdentifier = "E6"

type Identifier struct {
	*Int
}

func NewIdentifier() *Identifier {
	c := NewInt(TypeIdentifier)
	c.Format = FormatUInt32
	c.Permissions = []string{PermissionRead}
	c.SetMinValue(0)

	c.SetStepValue(1)
	c.SetValue(0)

	return &Identifier{c}
}
