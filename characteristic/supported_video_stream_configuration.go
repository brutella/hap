package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeSupportedVideoStreamConfiguration = "114"

type SupportedVideoStreamConfiguration struct {
	*Bytes
}

func NewSupportedVideoStreamConfiguration() *SupportedVideoStreamConfiguration {
	c := NewBytes(TypeSupportedVideoStreamConfiguration)
	c.Format = FormatTLV8
	c.Permissions = []string{PermissionRead}

	c.SetValue([]byte{})

	return &SupportedVideoStreamConfiguration{c}
}
