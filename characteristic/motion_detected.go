package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeMotionDetected = "22"

type MotionDetected struct {
	*Bool
}

func NewMotionDetected() *MotionDetected {
	c := NewBool(TypeMotionDetected)
	c.Format = FormatBool
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(false)

	return &MotionDetected{c}
}
