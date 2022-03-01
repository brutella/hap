package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	InputSourceTypeOther          int = 0
	InputSourceTypeHomeScreen     int = 1
	InputSourceTypeApplication    int = 10
	InputSourceTypeTuner          int = 2
	InputSourceTypeHdmi           int = 3
	InputSourceTypeCompositeVideo int = 4
	InputSourceTypeSVideo         int = 5
	InputSourceTypeComponentVideo int = 6
	InputSourceTypeDvi            int = 7
	InputSourceTypeAirplay        int = 8
	InputSourceTypeUsb            int = 9
)

const TypeInputSourceType = "DB"

type InputSourceType struct {
	*Int
}

func NewInputSourceType() *InputSourceType {
	c := NewInt(TypeInputSourceType)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(10)
	c.SetStepValue(1)
	c.SetValue(0)

	return &InputSourceType{c}
}
