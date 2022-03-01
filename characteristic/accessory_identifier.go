package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeAccessoryIdentifier = "57"

type AccessoryIdentifier struct {
	*String
}

func NewAccessoryIdentifier() *AccessoryIdentifier {
	c := NewString(TypeAccessoryIdentifier)
	c.Format = FormatString
	c.Permissions = []string{PermissionRead}
	c.Val = ""

	return &AccessoryIdentifier{c}
}
