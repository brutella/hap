package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	PictureModeOther          int = 0
	PictureModeStandard       int = 1
	PictureModeCalibrated     int = 2
	PictureModeCalibratedDark int = 3
	PictureModeVivid          int = 4
	PictureModeGame           int = 5
	PictureModeComputer       int = 6
	PictureModeCustom         int = 7
)

const TypePictureMode = "E2"

type PictureMode struct {
	*Int
}

func NewPictureMode() *PictureMode {
	c := NewInt(TypePictureMode)
	c.Format = FormatUInt16
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(13)
	c.SetStepValue(1)
	c.SetValue(0)

	return &PictureMode{c}
}
