package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	FilterChangeIndicationFilterOK     int = 0
	FilterChangeIndicationChangeFilter int = 1
)

const TypeFilterChangeIndication = "AC"

type FilterChangeIndication struct {
	*Int
}

func NewFilterChangeIndication() *FilterChangeIndication {
	c := NewInt(TypeFilterChangeIndication)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(0)

	return &FilterChangeIndication{c}
}
