package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	CurrentVisibilityStateShown  int = 0
	CurrentVisibilityStateHidden int = 1
)

const TypeCurrentVisibilityState = "135"

type CurrentVisibilityState struct {
	*Int
}

func NewCurrentVisibilityState() *CurrentVisibilityState {
	c := NewInt(TypeCurrentVisibilityState)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(3)
	c.SetStepValue(1)
	c.SetValue(0)

	return &CurrentVisibilityState{c}
}
