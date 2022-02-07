// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeObstructionDetected = "24"

type ObstructionDetected struct {
	*Bool
}

func NewObstructionDetected() *ObstructionDetected {
	c := NewBool(TypeObstructionDetected)
	c.Format = FormatBool
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(false)

	return &ObstructionDetected{c}
}
