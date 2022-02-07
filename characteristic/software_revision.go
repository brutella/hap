// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeSoftwareRevision = "54"

type SoftwareRevision struct {
	*String
}

func NewSoftwareRevision() *SoftwareRevision {
	c := NewString(TypeSoftwareRevision)
	c.Format = FormatString
	c.Permissions = []string{PermissionRead}
	c.Val = ""

	return &SoftwareRevision{c}
}
