package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeMute = "11A"

type Mute struct {
	*Bool
}

func NewMute() *Mute {
	c := NewBool(TypeMute)
	c.Format = FormatBool
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}

	c.SetValue(false)

	return &Mute{c}
}
