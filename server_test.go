package hap

import (
	"github.com/brutella/hap/accessory"

	"reflect"
	"testing"
)

func TestTxtRecords(t *testing.T) {
	expect := map[string]string{
		"pv": "1.0",
		"id": "1234",
		"c#": "1",
		"s#": "1",
		"sf": "0",
		"ff": "0",
		"md": "My MDNS Service",
		"ci": "1",
		"sh": "1ARVnw==",
	}

	a := accessory.New(accessory.Info{}, 1)
	a.Info.Name.SetValue("My MDNS Service")
	s, _ := NewServer(NewMemStore(), a)
	s.uuid = "1234"
	s.version = 1
	s.Protocol = "1.0"
	s.MfiCompliant = false

	if x := s.txtRecords(); reflect.DeepEqual(x, expect) == false {
		t.Fatalf("%v != %v", x, expect)
	}
}
