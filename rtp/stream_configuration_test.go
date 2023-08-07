package rtp

import (
	"fmt"
	"github.com/brutella/hap/characteristic"
	"github.com/brutella/hap/tlv8"
	"testing"
)

func TestSelectedStreamConfiguration(t *testing.T) {
	c := characteristic.NewSelectedStreamConfiguration()
	c.Val = "ARUCAQABEHW8tiJ9E0F4tLlvOURdFCc="

	b := c.Value()

	var cfg StreamConfiguration
	err := tlv8.Unmarshal(b, &cfg)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v", cfg)
}

func TestRTPParams(t *testing.T) {

	videoBuf := []byte{
		// tag, len, data
		1, 1, 99,
		2, 4, 26, 144, 146, 159,
		3, 2, 43, 1,
		4, 4, 0, 0, 0, 63,
		5, 2, 98, 5,
	}

	var videoParam RTPParams
	if err := tlv8.Unmarshal(videoBuf, &videoParam); err != nil {
		t.Error(err)
	}

	if videoParam.PayloadType != 99 {
		t.Error("Video PayloadType wrong:", videoParam.PayloadType)
	}

	if videoParam.Ssrc != 2677182490 {
		t.Error("Video Ssrc wrong", videoParam.Ssrc)
	}

	if videoParam.Bitrate != 299 {
		t.Error("Video Bitrate wrong", videoParam.Bitrate)
	}

	if videoParam.Interval != 0.5 {
		t.Error("Video Interval wrong", videoParam.Interval)
	}

	if videoParam.ComfortNoisePayloadType != 98 {
		t.Error("Video ComfortNoisePayloadType wrong", videoParam.ComfortNoisePayloadType)
	}

	audioBuf := []byte{
		// tag, len, data
		1, 1, 110,
		2, 4, 207, 83, 180, 9,
		3, 2, 24, 0,
		4, 4, 0, 0, 160, 64,
		6, 1, 13,
	}

	var audioParam RTPParams
	if err := tlv8.Unmarshal(audioBuf, &audioParam); err != nil {
		t.Error(err)
	}

	if audioParam.PayloadType != 110 {
		t.Error("Audio PayloadType wrong:", audioParam.PayloadType)
	}

	if audioParam.Ssrc != 162812879 {
		t.Error("Audio Ssrc wrong", audioParam.Ssrc)
	}

	if audioParam.Bitrate != 24 {
		t.Error("Audio Bitrate wrong", audioParam.Bitrate)
	}

	if audioParam.Interval != 5 {
		t.Error("Audio Interval wrong", audioParam.Interval)
	}

	if audioParam.MTU != 13 {
		t.Error("Audio MTU wrong", audioParam.MTU)
	}
}
