package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeRemainingDuration = "D4"

type RemainingDuration struct {
	*Int
}

func NewRemainingDuration() *RemainingDuration {
	c := NewInt(TypeRemainingDuration)
	c.Format = FormatUInt32
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(3600)
	c.SetStepValue(1)
	c.SetValue(0)

	return &RemainingDuration{c}
}
