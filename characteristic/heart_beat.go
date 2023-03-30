package characteristic

const TypeHeartBeat = "24A"

type HeartBeat struct {
	*Int
}

func NewHeartBeat() *HeartBeat {
	c := NewInt(TypeHeartBeat)
	c.Format = FormatUInt32
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetValue(0)

	return &HeartBeat{c}
}
