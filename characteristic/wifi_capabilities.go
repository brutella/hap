package characteristic

const TypeWifiCapabilities = "22C"

type WifiCapabilities = struct {
	*Int
}

func NewWifiCapabilities() *WifiCapabilities {
	c := NewInt(TypeWifiCapabilities)
	c.Format = FormatUInt32
	c.Permissions = []string{PermissionRead}
	c.Val = 1
	return &WifiCapabilities{c}
}
