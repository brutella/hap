package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	TargetMediaStatePlay  int = 0
	TargetMediaStatePause int = 1
	TargetMediaStateStop  int = 2
)

const TypeTargetMediaState = "137"

type TargetMediaState struct {
	*Int
}

func NewTargetMediaState() *TargetMediaState {
	c := NewInt(TypeTargetMediaState)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(2)
	c.SetStepValue(1)
	c.SetValue(0)

	return &TargetMediaState{c}
}
