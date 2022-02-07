// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeResetFilterIndication = "AD"

type ResetFilterIndication struct {
	*Int
}

func NewResetFilterIndication() *ResetFilterIndication {
	c := NewInt(TypeResetFilterIndication)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionWrite}
	c.SetMinValue(1)
	c.SetMaxValue(1)
	c.SetStepValue(1)

	return &ResetFilterIndication{c}
}
