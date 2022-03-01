package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	ServiceLabelNamespaceDots           int = 0
	ServiceLabelNamespaceArabicNumerals int = 1
)

const TypeServiceLabelNamespace = "CD"

type ServiceLabelNamespace struct {
	*Int
}

func NewServiceLabelNamespace() *ServiceLabelNamespace {
	c := NewInt(TypeServiceLabelNamespace)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead}

	c.SetValue(0)

	return &ServiceLabelNamespace{c}
}
