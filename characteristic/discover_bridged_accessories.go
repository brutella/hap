package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	DiscoverBridgedAccessoriesStartDiscovery int = 0
	DiscoverBridgedAccessoriesStopDiscovery  int = 1
)

const TypeDiscoverBridgedAccessories = "9E"

type DiscoverBridgedAccessories struct {
	*Int
}

func NewDiscoverBridgedAccessories() *DiscoverBridgedAccessories {
	c := NewInt(TypeDiscoverBridgedAccessories)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}
	c.Val = 0

	return &DiscoverBridgedAccessories{c}
}
