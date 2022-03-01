package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	VolumeSelectorIncrement int = 0
	VolumeSelectorDecrement int = 1
)

const TypeVolumeSelector = "EA"

type VolumeSelector struct {
	*Int
}

func NewVolumeSelector() *VolumeSelector {
	c := NewInt(TypeVolumeSelector)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionWrite}
	c.SetMinValue(0)
	c.SetMaxValue(1)
	c.SetStepValue(1)

	return &VolumeSelector{c}
}
