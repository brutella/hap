package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeCurrentPosition = "6D"

type CurrentPosition struct {
	*Int
}

func NewCurrentPosition() *CurrentPosition {
	c := NewInt(TypeCurrentPosition)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(100)
	c.SetStepValue(1)
	c.SetValue(0)
	c.Unit = UnitPercentage

	return &CurrentPosition{c}
}
