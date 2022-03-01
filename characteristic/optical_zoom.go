package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeOpticalZoom = "11C"

type OpticalZoom struct {
	*Float
}

func NewOpticalZoom() *OpticalZoom {
	c := NewFloat(TypeOpticalZoom)
	c.Format = FormatFloat
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}

	c.SetValue(0)

	return &OpticalZoom{c}
}
