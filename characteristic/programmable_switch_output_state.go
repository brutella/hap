package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeProgrammableSwitchOutputState = "74"

type ProgrammableSwitchOutputState struct {
	*Int
}

func NewProgrammableSwitchOutputState() *ProgrammableSwitchOutputState {
	c := NewInt(TypeProgrammableSwitchOutputState)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(1)
	c.SetStepValue(1)
	c.Val = 0

	return &ProgrammableSwitchOutputState{c}
}
