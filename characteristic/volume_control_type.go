package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	VolumeControlTypeNone                int = 0
	VolumeControlTypeRelative            int = 1
	VolumeControlTypeRelativeWithCurrent int = 2
	VolumeControlTypeAbsolute            int = 3
)

const TypeVolumeControlType = "E9"

type VolumeControlType struct {
	*Int
}

func NewVolumeControlType() *VolumeControlType {
	c := NewInt(TypeVolumeControlType)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(3)
	c.SetStepValue(1)
	c.SetValue(0)

	return &VolumeControlType{c}
}
