package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	InputDeviceTypeOther       int = 0
	InputDeviceTypeTv          int = 1
	InputDeviceTypeRecording   int = 2
	InputDeviceTypeTuner       int = 3
	InputDeviceTypePlayback    int = 4
	InputDeviceTypeAudioSystem int = 5
)

const TypeInputDeviceType = "DC"

type InputDeviceType struct {
	*Int
}

func NewInputDeviceType() *InputDeviceType {
	c := NewInt(TypeInputDeviceType)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(5)
	c.SetStepValue(1)
	c.SetValue(0)

	return &InputDeviceType{c}
}
