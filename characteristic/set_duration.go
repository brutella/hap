package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeSetDuration = "D3"

type SetDuration struct {
	*Int
}

func NewSetDuration() *SetDuration {
	c := NewInt(TypeSetDuration)
	c.Format = FormatUInt32
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(3600)
	c.SetStepValue(1)
	c.SetValue(0)

	return &SetDuration{c}
}
