package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeOutletInUse = "26"

type OutletInUse struct {
	*Bool
}

func NewOutletInUse() *OutletInUse {
	c := NewBool(TypeOutletInUse)
	c.Format = FormatBool
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(false)

	return &OutletInUse{c}
}
