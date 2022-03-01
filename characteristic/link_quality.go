package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeLinkQuality = "9C"

type LinkQuality struct {
	*Int
}

func NewLinkQuality() *LinkQuality {
	c := NewInt(TypeLinkQuality)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(1)
	c.SetMaxValue(4)
	c.SetStepValue(1)
	c.Val = 1

	return &LinkQuality{c}
}
