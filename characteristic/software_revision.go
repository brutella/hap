package characteristic

// THIS FILE IS AUTO-GENERATED

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
