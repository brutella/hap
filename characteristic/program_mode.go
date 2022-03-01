package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	ProgramModeNoProgramScheduled         int = 0
	ProgramModeProgramScheduled           int = 1
	ProgramModeProgramScheduledManualMode int = 2
)

const TypeProgramMode = "D1"

type ProgramMode struct {
	*Int
}

func NewProgramMode() *ProgramMode {
	c := NewInt(TypeProgramMode)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(0)

	return &ProgramMode{c}
}
