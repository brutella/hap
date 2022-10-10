package characteristic

const TypeActivityInterval = "23B"

type ActivityInterval struct {
	*Int
}

func NewActivityInterval() *ActivityInterval {
	c := NewInt(TypeActivityInterval)
	c.Format = FormatUInt32
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(0)
	c.SetStepValue(1)
	c.SetValue(0)

	return &ActivityInterval{c}
}
