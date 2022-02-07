// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hap/characteristic"
)

const TypeCameraRecordingManagement = "204"

type CameraRecordingManagement struct {
	*S

	SupportedCameraRecordingConfiguration *characteristic.SupportedCameraRecordingConfiguration
	SupportedVideoRecordingConfiguration  *characteristic.SupportedVideoRecordingConfiguration
	SupportedAudioRecordingConfiguration  *characteristic.SupportedAudioRecordingConfiguration
	SelectedCameraRecordingConfiguration  *characteristic.SelectedCameraRecordingConfiguration
}

func NewCameraRecordingManagement() *CameraRecordingManagement {
	s := CameraRecordingManagement{}
	s.S = New(TypeCameraRecordingManagement)

	s.SupportedCameraRecordingConfiguration = characteristic.NewSupportedCameraRecordingConfiguration()
	s.AddC(s.SupportedCameraRecordingConfiguration.C)

	s.SupportedVideoRecordingConfiguration = characteristic.NewSupportedVideoRecordingConfiguration()
	s.AddC(s.SupportedVideoRecordingConfiguration.C)

	s.SupportedAudioRecordingConfiguration = characteristic.NewSupportedAudioRecordingConfiguration()
	s.AddC(s.SupportedAudioRecordingConfiguration.C)

	s.SelectedCameraRecordingConfiguration = characteristic.NewSelectedCameraRecordingConfiguration()
	s.AddC(s.SelectedCameraRecordingConfiguration.C)

	return &s
}
