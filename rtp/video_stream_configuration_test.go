package rtp

import (
	"reflect"
	"testing"

	"github.com/brutella/hap/tlv8"
)

func TestMarhsalUnmarshalDefaultVideoStreamConfiguration(t *testing.T) {
	want := DefaultVideoStreamConfiguration()
	buf, err := tlv8.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}

	var is VideoStreamConfiguration
	err = tlv8.Unmarshal(buf, &is)
	if err != nil {
		t.Fatal(err)
	}

	if reflect.DeepEqual(is, want) == false {
		t.Fatalf("is=%+v want=%+v", is, want)
	}
}
