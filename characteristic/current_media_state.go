package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	CurrentMediaStatePlay    int = 0
	CurrentMediaStatePause   int = 1
	CurrentMediaStateStop    int = 2
	CurrentMediaStateUnknown int = 3
)

const TypeCurrentMediaState = "E0"

type CurrentMediaState struct {
	*Int
}

func NewCurrentMediaState() *CurrentMediaState {
	c := NewInt(TypeCurrentMediaState)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(3)
	c.SetStepValue(1)
	c.SetValue(0)
	c.Unit = UnitPercentage

	return &CurrentMediaState{c}
}
