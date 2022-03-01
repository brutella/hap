package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeNightVision = "11B"

type NightVision struct {
	*Bool
}

func NewNightVision() *NightVision {
	c := NewBool(TypeNightVision)
	c.Format = FormatBool
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}

	c.SetValue(false)

	return &NightVision{c}
}
