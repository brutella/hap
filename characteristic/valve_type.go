package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	ValveTypeGenericValve int = 0
	ValveTypeIrrigation   int = 1
	ValveTypeShowerHead   int = 2
	ValveTypeWaterFaucet  int = 3
)

const TypeValveType = "D5"

type ValveType struct {
	*Int
}

func NewValveType() *ValveType {
	c := NewInt(TypeValveType)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(0)

	return &ValveType{c}
}
