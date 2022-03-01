package characteristic

// THIS FILE IS AUTO-GENERATED

const (
	CurrentHeaterCoolerStateInactive int = 0
	CurrentHeaterCoolerStateIdle     int = 1
	CurrentHeaterCoolerStateHeating  int = 2
	CurrentHeaterCoolerStateCooling  int = 3
)

const TypeCurrentHeaterCoolerState = "B1"

type CurrentHeaterCoolerState struct {
	*Int
}

func NewCurrentHeaterCoolerState() *CurrentHeaterCoolerState {
	c := NewInt(TypeCurrentHeaterCoolerState)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionEvents}

	c.SetValue(0)

	return &CurrentHeaterCoolerState{c}
}
