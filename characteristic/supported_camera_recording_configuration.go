package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeSupportedCameraRecordingConfiguration = "205"

type SupportedCameraRecordingConfiguration struct {
	*Bytes
}

func NewSupportedCameraRecordingConfiguration() *SupportedCameraRecordingConfiguration {
	c := NewBytes(TypeSupportedCameraRecordingConfiguration)
	c.Format = FormatTLV8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue([]byte{})

	return &SupportedCameraRecordingConfiguration{c}
}
