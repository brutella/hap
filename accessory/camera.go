package accessory

import (
	"github.com/brutella/hap/service"
)

// Camera provides RTP video streaming.
type Camera struct {
	*A
	Control           *service.CameraControl
	StreamManagement1 *service.CameraRTPStreamManagement
	StreamManagement2 *service.CameraRTPStreamManagement
}

// NewCamera returns an IP camera accessory.
func NewCamera(info Info) *Camera {
	a := Camera{}
	a.A = New(info, TypeIPCamera)
	
	a.Control = service.NewCameraControl()
	a.AddS(a.Control.S)

	// TODO (mah) a camera must support at least 2 rtp streams
	a.StreamManagement1 = service.NewCameraRTPStreamManagement()
	a.StreamManagement2 = service.NewCameraRTPStreamManagement()
	a.AddS(a.StreamManagement1.S)
	// a.AddS(a.StreamManagement2.S)

	return &a
}
