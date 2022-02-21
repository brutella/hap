package accessory

type Bridge struct {
	*A
}

// NewBridge returns a bridge which implements model.Bridge.
func NewBridge(info Info) *Bridge {
	a := Bridge{}
	a.A = New(info, TypeBridge)

	return &a
}
