package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeImageMirroring = "11F"

type ImageMirroring struct {
	*Bool
}

func NewImageMirroring() *ImageMirroring {
	c := NewBool(TypeImageMirroring)
	c.Format = FormatBool
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}

	c.SetValue(false)

	return &ImageMirroring{c}
}
