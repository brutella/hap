package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeVolume = "119"

type Volume struct {
	*Int
}

func NewVolume() *Volume {
	c := NewInt(TypeVolume)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(100)
	c.SetStepValue(1)
	c.SetValue(0)
	c.Unit = UnitPercentage

	return &Volume{c}
}
