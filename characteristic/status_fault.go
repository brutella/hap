package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	StatusFaultNoFault      int = 0
	StatusFaultGeneralFault int = 1
)

const TypeStatusFault = "77"

type StatusFault struct {
	*Int
}

func NewStatusFault() *StatusFault {
	c := NewInt(TypeStatusFault)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(0)

	return &StatusFault{c}
}
