package rtp

import (
	"testing"

	"github.com/brutella/hap/characteristic"
	"github.com/brutella/hap/tlv8"
)

func TestStreamController(t *testing.T) {
	c := characteristic.NewSupportedVideoStreamConfiguration()
	c.Val = "AX8BAQACDAEBAQEBAgIBAAMBAAECgAcCAjgEAwEeAQIABQIC0AIDAR4BAoACAgJoAQMBHgEC4AECAg4BAwEeAQJAAQICtAADAR4BAgAFAgLAAwMBHgECAAQCAgADAwEeAQKAAgIC4AEDAR4BAuABAgJoAQMBHgECQAECAvAAAwEP"

	b := c.Value()
	if len(b) == 0 {
		t.Fatal("Zero length bytes")
	}

	var cfg VideoStreamConfiguration
	err := tlv8.Unmarshal(b, &cfg)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMarshalVideoCodecConfiguration(t *testing.T) {
	codec := NewH264VideoCodecConfiguration()
	b, err := tlv8.Marshal(codec)
	if err != nil {
		t.Fatal(err)
	}

	var c VideoCodecConfiguration
	err = tlv8.Unmarshal(b, &c)
	if err != nil {
		t.Fatal(err)
	}
}

func TestStreamingStatus(t *testing.T) {
	c := characteristic.NewStreamingStatus()
	c.Val = "AQEA"

	b := c.Value()
	if len(b) == 0 {
		t.Fatalf("Zero length bytes")
	}

	var status StreamingStatus
	err := tlv8.Unmarshal(b, &status)
	if err != nil {
		t.Fatal(err)
	}
}
