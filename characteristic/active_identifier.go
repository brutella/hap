package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeActiveIdentifier = "E7"

type ActiveIdentifier struct {
	*Int
}

func NewActiveIdentifier() *ActiveIdentifier {
	c := NewInt(TypeActiveIdentifier)
	c.Format = FormatUInt32
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}
	c.SetMinValue(0)

	c.SetValue(0)

	return &ActiveIdentifier{c}
}
