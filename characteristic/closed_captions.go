package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	ClosedCaptionsDisabled int = 0
	ClosedCaptionsEnabled  int = 1
)

const TypeClosedCaptions = "DD"

type ClosedCaptions struct {
	*Int
}

func NewClosedCaptions() *ClosedCaptions {
	c := NewInt(TypeClosedCaptions)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(1)
	c.SetStepValue(1)
	c.SetValue(0)

	return &ClosedCaptions{c}
}
