package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeCategory = "A3"

type Category struct {
	*Int
}

func NewCategory() *Category {
	c := NewInt(TypeCategory)
	c.Format = FormatUInt16
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(1)
	c.SetMaxValue(16)
	c.SetStepValue(1)
	c.Val = 1

	return &Category{c}
}
