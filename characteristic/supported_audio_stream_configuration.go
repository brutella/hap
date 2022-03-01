package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeSupportedAudioStreamConfiguration = "115"

type SupportedAudioStreamConfiguration struct {
	*Bytes
}

func NewSupportedAudioStreamConfiguration() *SupportedAudioStreamConfiguration {
	c := NewBytes(TypeSupportedAudioStreamConfiguration)
	c.Format = FormatTLV8
	c.Permissions = []string{PermissionRead}

	c.SetValue([]byte{})

	return &SupportedAudioStreamConfiguration{c}
}
