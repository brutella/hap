// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeHardwareRevision = "53"

type HardwareRevision struct {
	*String
}

func NewHardwareRevision() *HardwareRevision {
	c := NewString(TypeHardwareRevision)
	c.Format = FormatString
	c.Permissions = []string{PermissionRead}

	c.SetValue("")

	return &HardwareRevision{c}
}
