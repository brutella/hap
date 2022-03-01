package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeSerialNumber = "30"

type SerialNumber struct {
	*String
}

func NewSerialNumber() *SerialNumber {
	c := NewString(TypeSerialNumber)
	c.Format = FormatString
	c.Permissions = []string{PermissionRead}

	c.SetValue("")

	return &SerialNumber{c}
}
