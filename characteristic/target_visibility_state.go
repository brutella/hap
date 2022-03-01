package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	TargetVisibilityStateShown  int = 0
	TargetVisibilityStateHidden int = 1
)

const TypeTargetVisibilityState = "134"

type TargetVisibilityState struct {
	*Int
}

func NewTargetVisibilityState() *TargetVisibilityState {
	c := NewInt(TypeTargetVisibilityState)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(2)
	c.SetStepValue(1)
	c.SetValue(0)

	return &TargetVisibilityState{c}
}
