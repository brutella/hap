package characteristic

const TypeWifiConfigurationControl = "22D"

type WifiConfigurationControl = struct {
	*Bytes
}

func NewWifiConfigurationControl() *WifiConfigurationControl {
	c := NewBytes(TypeWifiCapabilities)
	c.Format = FormatTLV8
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}

	return &WifiConfigurationControl{c}
}
