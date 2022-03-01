package characteristic

// THIS FILE IS AUTO-GENERATED

const TypeSecuritySystemAlarmType = "8E"

type SecuritySystemAlarmType struct {
	*Int
}

func NewSecuritySystemAlarmType() *SecuritySystemAlarmType {
	c := NewInt(TypeSecuritySystemAlarmType)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}
	c.SetMinValue(0)
	c.SetMaxValue(1)
	c.SetStepValue(1)
	c.SetValue(0)

	return &SecuritySystemAlarmType{c}
}
