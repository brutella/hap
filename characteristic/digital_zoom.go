package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeDigitalZoom = "11D"

type DigitalZoom struct {
	*Float
}

func NewDigitalZoom() *DigitalZoom {
	c := NewFloat(TypeDigitalZoom)
	c.Format = FormatFloat
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}

	c.SetValue(0)

	return &DigitalZoom{c}
}
