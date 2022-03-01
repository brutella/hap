package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	PowerModeSelectionShow int = 0
	PowerModeSelectionHide int = 1
)

const TypePowerModeSelection = "DF"

type PowerModeSelection struct {
	*Int
}

func NewPowerModeSelection() *PowerModeSelection {
	c := NewInt(TypePowerModeSelection)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionWrite}
	c.SetMinValue(0)
	c.SetMaxValue(1)
	c.SetStepValue(1)

	return &PowerModeSelection{c}
}
