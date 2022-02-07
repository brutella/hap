// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeDayOfTheWeek = "98"

type DayOfTheWeek struct {
	*Int
}

func NewDayOfTheWeek() *DayOfTheWeek {
	c := NewInt(TypeDayOfTheWeek)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionWrite}
	c.SetMinValue(1)
	c.SetMaxValue(7)
	c.SetStepValue(1)
	c.Val = 1

	return &DayOfTheWeek{c}
}
