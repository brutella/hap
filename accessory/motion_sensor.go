package accessory

import "github.com/brutella/hap/service"

type MotionSensor struct {
	*A
	MotionSensor *service.MotionSensor
}

// NewMotionSensor returns a motion sensor.
func NewMotionSensor(info Info) *MotionSensor {
	a := MotionSensor{}
	a.A = New(info, TypeSensor)

	a.MotionSensor = service.NewMotionSensor()
	a.AddS(a.MotionSensor.S)

	return &a
}
