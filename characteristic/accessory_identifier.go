// THIS FILE IS AUTO-GENERATED
package characteristic

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
