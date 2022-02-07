// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	ProgrammableSwitchEventSinglePress int = 0
	ProgrammableSwitchEventDoublePress int = 1
	ProgrammableSwitchEventLongPress   int = 2
)

const TypeProgrammableSwitchEvent = "73"

type ProgrammableSwitchEvent struct {
	*Int
}

func NewProgrammableSwitchEvent() *ProgrammableSwitchEvent {
	c := NewInt(TypeProgrammableSwitchEvent)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(0)
	c.updateOnSameValue = true

	return &ProgrammableSwitchEvent{c}
}
