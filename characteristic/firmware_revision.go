package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeFirmwareRevision = "52"

type FirmwareRevision struct {
	*String
}

func NewFirmwareRevision() *FirmwareRevision {
	c := NewString(TypeFirmwareRevision)
	c.Format = FormatString
	c.Permissions = []string{PermissionRead}

	c.SetValue("")

	return &FirmwareRevision{c}
}
