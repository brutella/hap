package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeAccessoryFlags = "A6"

type AccessoryFlags struct {
	*Int
}

func NewAccessoryFlags() *AccessoryFlags {
	c := NewInt(TypeAccessoryFlags)
	c.Format = FormatUInt32
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(0)

	return &AccessoryFlags{c}
}
