package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeConfigureBridgedAccessoryStatus = "9D"

type ConfigureBridgedAccessoryStatus struct {
	*Bytes
}

func NewConfigureBridgedAccessoryStatus() *ConfigureBridgedAccessoryStatus {
	c := NewBytes(TypeConfigureBridgedAccessoryStatus)
	c.Format = FormatTLV8
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.Val = []byte{}

	return &ConfigureBridgedAccessoryStatus{c}
}
