// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeCameraRTPStreamManagement = "110"

type CameraRTPStreamManagement struct {
	*S

	SupportedVideoStreamConfiguration *characteristic.SupportedVideoStreamConfiguration
	SupportedAudioStreamConfiguration *characteristic.SupportedAudioStreamConfiguration
	SupportedRTPConfiguration         *characteristic.SupportedRTPConfiguration
	SelectedRTPStreamConfiguration    *characteristic.SelectedRTPStreamConfiguration
	StreamingStatus                   *characteristic.StreamingStatus
	SetupEndpoints                    *characteristic.SetupEndpoints
}

func NewCameraRTPStreamManagement() *CameraRTPStreamManagement {
	s := CameraRTPStreamManagement{}
	s.S = New(TypeCameraRTPStreamManagement)

	s.SupportedVideoStreamConfiguration = characteristic.NewSupportedVideoStreamConfiguration()
	s.AddC(s.SupportedVideoStreamConfiguration.C)

	s.SupportedAudioStreamConfiguration = characteristic.NewSupportedAudioStreamConfiguration()
	s.AddC(s.SupportedAudioStreamConfiguration.C)

	s.SupportedRTPConfiguration = characteristic.NewSupportedRTPConfiguration()
	s.AddC(s.SupportedRTPConfiguration.C)

	s.SelectedRTPStreamConfiguration = characteristic.NewSelectedRTPStreamConfiguration()
	s.AddC(s.SelectedRTPStreamConfiguration.C)

	s.StreamingStatus = characteristic.NewStreamingStatus()
	s.AddC(s.StreamingStatus.C)

	s.SetupEndpoints = characteristic.NewSetupEndpoints()
	s.AddC(s.SetupEndpoints.C)

	return &s
}
