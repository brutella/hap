// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeSelectedCameraRecordingConfiguration = "209"

type SelectedCameraRecordingConfiguration struct {
	*Bytes
}

func NewSelectedCameraRecordingConfiguration() *SelectedCameraRecordingConfiguration {
	c := NewBytes(TypeSelectedCameraRecordingConfiguration)
	c.Format = FormatTLV8
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}

	c.SetValue([]byte{})

	return &SelectedCameraRecordingConfiguration{c}
}
