// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeHoldPosition = "6F"

type HoldPosition struct {
	*Bool
}

func NewHoldPosition() *HoldPosition {
	c := NewBool(TypeHoldPosition)
	c.Format = FormatBool
	c.Permissions = []string{PermissionWrite}

	return &HoldPosition{c}
}
