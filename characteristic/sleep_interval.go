package characteristic

const TypeSleepInterval = "23A"

type SleepInterval struct {
	*Int
}

func NewSleepInterval() *SleepInterval {
	c := NewInt(TypeSleepInterval)
	c.Format = FormatUInt32
	c.Permissions = []string{PermissionRead}
	c.SetMinValue(0)
	c.SetMaxValue(67108863)
	c.SetStepValue(1)
	c.SetValue(0)

	return &SleepInterval{c}
}
