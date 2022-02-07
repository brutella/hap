// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeSupportedAudioRecordingConfiguration = "207"

type SupportedAudioRecordingConfiguration struct {
	*Bytes
}

func NewSupportedAudioRecordingConfiguration() *SupportedAudioRecordingConfiguration {
	c := NewBytes(TypeSupportedAudioRecordingConfiguration)
	c.Format = FormatTLV8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue([]byte{})

	return &SupportedAudioRecordingConfiguration{c}
}
