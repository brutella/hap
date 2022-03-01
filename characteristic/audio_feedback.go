package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeAudioFeedback = "5"

type AudioFeedback struct {
	*Bool
}

func NewAudioFeedback() *AudioFeedback {
	c := NewBool(TypeAudioFeedback)
	c.Format = FormatBool
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}

	c.SetValue(false)

	return &AudioFeedback{c}
}
