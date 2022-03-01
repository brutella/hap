package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	SwingModeSwingDisabled int = 0
	SwingModeSwingEnabled  int = 1
)

const TypeSwingMode = "B6"

type SwingMode struct {
	*Int
}

func NewSwingMode() *SwingMode {
	c := NewInt(TypeSwingMode)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}

	c.SetValue(0)

	return &SwingMode{c}
}
