package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	CurrentSlatStateFixed    int = 0
	CurrentSlatStateJammed   int = 1
	CurrentSlatStateSwinging int = 2
)

const TypeCurrentSlatState = "AA"

type CurrentSlatState struct {
	*Int
}

func NewCurrentSlatState() *CurrentSlatState {
	c := NewInt(TypeCurrentSlatState)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(0)

	return &CurrentSlatState{c}
}
