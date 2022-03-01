package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeManufacturer = "20"

type Manufacturer struct {
	*String
}

func NewManufacturer() *Manufacturer {
	c := NewString(TypeManufacturer)
	c.Format = FormatString
	c.Permissions = []string{PermissionRead}

	c.SetValue("")

	return &Manufacturer{c}
}
