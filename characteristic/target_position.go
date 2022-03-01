package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeTargetPosition = "7C"

type TargetPosition struct {
	*Int
}

func NewTargetPosition() *TargetPosition {
	c := NewInt(TypeTargetPosition)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(100)
	c.SetStepValue(1)
	c.SetValue(0)
	c.Unit = UnitPercentage

	return &TargetPosition{c}
}
