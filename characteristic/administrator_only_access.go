// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeAdministratorOnlyAccess = "1"

type AdministratorOnlyAccess struct {
	*Bool
}

func NewAdministratorOnlyAccess() *AdministratorOnlyAccess {
	c := NewBool(TypeAdministratorOnlyAccess)
	c.Format = FormatBool
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}

	c.SetValue(false)

	return &AdministratorOnlyAccess{c}
}
