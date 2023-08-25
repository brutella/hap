package characteristic

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestCharacteristicSetValue(t *testing.T) {
	req := &http.Request{}
	c := NewBrightness()
	c.Val = 0

	n := 0
	c.OnValueUpdate(func(new, old int, r *http.Request) {
		if r != req {
			t.Fatal(r)
		}
		n++
	})

	c.SetValueRequest(10, req)
	if is, want := c.Value(), 10; is != want {
		t.Fatalf("%v != %v", is, want)
	}

	c.SetValueRequest(20, req)
	if is, want := c.Value(), 20; is != want {
		t.Fatalf("%v != %v", is, want)
	}

	if is, want := n, 2; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestCharacteristicValueTypeConversion(t *testing.T) {
	c := NewBrightness()
	c.Val = 5
	c.setValue(float64(20.5), nil)

	if is, want := c.Val, 20; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	c.setValue("91", nil)

	if is, want := c.Val, 91; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	c.setValue(true, nil)

	if is, want := c.Val, 1; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestCharacteristicOnValueUpdate(t *testing.T) {
	c := NewBrightness()
	c.Val = 5

	d := false
	c.OnValueUpdate(func(new, old int, r *http.Request) {
		if r != nil {
			t.Fatal(r)
		}

		if is, want := old, 5; is != want {
			t.Fatalf("%v != %v", is, want)
		}

		if is, want := new, 6; is != want {
			t.Fatalf("%v != %v", is, want)
		}
		d = true
	})

	c.SetValue(6)

	if is, want := d, true; is != want {
		t.Fatalf("%v != %v", is, want)
	}
}

func TestValueChange(t *testing.T) {
	c := NewProgrammableSwitchEvent()
	c.Val = ProgrammableSwitchEventSinglePress

	changed := false
	c.OnValueUpdate(func(new, old int, r *http.Request) {
		changed = true
	})

	c.SetValue(ProgrammableSwitchEventSinglePress)

	if is, want := changed, true; is != want {
		t.Fatalf("%v != %v", is, want)
	}
}

func TestValueIngoreValueUpdate(t *testing.T) {
	c := NewBrightness()
	c.Val = 5

	c.OnValueUpdate(func(new, old int, r *http.Request) {
		t.Fatalf("Update value from %v to %v is unexpected", old, new)
	})

	c.SetValue(5)
}

func TestReadOnly(t *testing.T) {
	c := NewName()

	c.SetValue("Matthias")
	c.SetValueRequest("Gottfried", &http.Request{})

	if is, want := c.Value(), "Matthias"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	c.SetValueRequest("Gottfried", nil)
	if is, want := c.Value(), "Gottfried"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestSetValueRequestFuncError(t *testing.T) {
	c := NewBrightness()

	c.SetValue(100)
	c.SetValueRequestFunc = func(v interface{}, r *http.Request) (response interface{}, status int) {
		if r != nil {
			status = -70408
		}

		return
	}

	_, s := c.SetValueRequest(50, &http.Request{})
	if is, want := s, -70408; is != want {
		t.Fatalf("%v != %v", is, want)
	}

	if is, want := c.Value(), 100; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestOnSetRemoteValue(t *testing.T) {
	c := NewBrightness()

	c.SetValue(100)
	c.OnSetRemoteValue(func(v int) error {
		return nil
	})

	v, s := c.SetValueRequest(50, &http.Request{})
	if is, want := s, 0; is != want {
		t.Fatalf("%v != %v", is, want)
	}

	if is, want := v, 50; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestValidValues(t *testing.T) {
	c := NewTargetHeaterCoolerState()
	c.ValidVals = []int{TargetHeaterCoolerStateAuto, TargetHeaterCoolerStateHeat}

	if err := c.SetValue(TargetHeaterCoolerStateCool); err == nil {
		t.Fatal("invalid value error expected")
	}

	if err := c.SetValue(TargetHeaterCoolerStateHeat); err != nil {
		t.Fatal("no error expected")
	}
}

func TestValidRange(t *testing.T) {
	c := NewTargetHeaterCoolerState()
	c.ValidRange = []int{TargetHeaterCoolerStateAuto, TargetHeaterCoolerStateHeat}

	if err := c.SetValue(TargetHeaterCoolerStateCool); err == nil {
		t.Fatal("invalid value error expected")
	}

	if err := c.SetValue(TargetHeaterCoolerStateHeat); err != nil {
		t.Fatal("no error expected")
	}
}

func encodeDecodeJson(c *C, t *testing.T) map[string]interface{} {
	j, err := c.MarshalJSON()
	if err != nil {
		t.Fatal("cannot MarshalJSON: ", err)
	}

	var jsonMap map[string]interface{}
	err = json.Unmarshal(j, &jsonMap)
	if err != nil {
		t.Fatal("invalid encoded JSON: ", err)
	}

	return jsonMap
}

func TestCharacteristicJson(t *testing.T) {
	cs := []*C{
		//NewContactSensorState().C,  // int
		NewCurrentTemperature().C,  // float
		NewCurrentTransport().C,    // bool
		NewAccessoryIdentifier().C, // string
	}

	for _, c := range cs {
		jsonMap := encodeDecodeJson(c, t)

		// verify properties
		if is, want := jsonMap["type"], c.Type; is != want {
			t.Fatalf("marshaled type is wrong: is=%v wanted=%v", is, want)
		}
		if is, want := jsonMap["format"], c.Format; is != want {
			t.Fatalf("marshaled format is wrong: is=%v wanted=%v", is, want)
		}
		if is, want := jsonMap["value"], c.Val; is != want {
			t.Fatalf("marshaled value is wrong: is=%v wanted=%v", is, want)
		}

		// set a ValueRequestFunc that returns an error
		c.ValueRequestFunc = func(r *http.Request) (response interface{}, status int) {
			return nil, -70408
		}

		// re-encode
		jsonMap = encodeDecodeJson(c, t)

		jv, exists := jsonMap["value"]
		if !exists {
			t.Fatalf("errored characteristic is missing \"value\": %+v", jsonMap)
		}
		if is, want := reflect.TypeOf(jv), reflect.TypeOf(c.Val); is != want {
			t.Fatalf("json-encoded value is of wrong type: is=%v want=%v", is, want)
		}
	}

	// special case /identify must not emit any "value"
	id := NewIdentify().C
	jsonMap := encodeDecodeJson(id, t)

	if _, exists := jsonMap["value"]; exists {
		t.Fatalf("Identify characteristic cannot emit \"value\": %+v", jsonMap)
	}
}
