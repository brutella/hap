package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeSupportedRTPConfiguration = "116"

type SupportedRTPConfiguration struct {
	*Bytes
}

func NewSupportedRTPConfiguration() *SupportedRTPConfiguration {
	c := NewBytes(TypeSupportedRTPConfiguration)
	c.Format = FormatTLV8
	c.Permissions = []string{PermissionRead}

	c.SetValue([]byte{})

	return &SupportedRTPConfiguration{c}
}
