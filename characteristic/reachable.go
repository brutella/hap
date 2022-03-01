package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeReachable = "63"

type Reachable struct {
	*Bool
}

func NewReachable() *Reachable {
	c := NewBool(TypeReachable)
	c.Format = FormatBool
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.Val = false

	return &Reachable{c}
}
