// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	TemperatureDisplayUnitsCelsius    int = 0
	TemperatureDisplayUnitsFahrenheit int = 1
)

const TypeTemperatureDisplayUnits = "36"

type TemperatureDisplayUnits struct {
	*Int
}

func NewTemperatureDisplayUnits() *TemperatureDisplayUnits {
	c := NewInt(TypeTemperatureDisplayUnits)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}

	c.SetValue(0)

	return &TemperatureDisplayUnits{c}
}
