package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeDiscoveredBridgedAccessories = "9F"

type DiscoveredBridgedAccessories struct {
	*Int
}

func NewDiscoveredBridgedAccessories() *DiscoveredBridgedAccessories {
	c := NewInt(TypeDiscoveredBridgedAccessories)
	c.Format = FormatUInt16
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.Val = 0

	return &DiscoveredBridgedAccessories{c}
}
