package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeConfigureBridgedAccessory = "A0"

type ConfigureBridgedAccessory struct {
	*Bytes
}

func NewConfigureBridgedAccessory() *ConfigureBridgedAccessory {
	c := NewBytes(TypeConfigureBridgedAccessory)
	c.Format = FormatTLV8
	c.Permissions = []string{PermissionWrite}

	return &ConfigureBridgedAccessory{c}
}
