package characteristic

import (
	"sync"

	"github.com/brutella/hap/log"
	"github.com/xiam/to"

	"encoding/json"
	"net/http"
)

const (
	PermissionRead          = "pr" // The characteristic can only be read by paired controllers.
	PermissionWrite         = "pw" // The characteristic can only be written by paired controllers.
	PermissionTimedWrite    = "tw" // The characteristic allows only timed write procedure.
	PermissionEvents        = "ev" // The characteristic supports events.
	PermissionHidden        = "hd" // The characteristic is hidden from the user.
	PermissionWriteResponse = "wr" // The characteristic supports write response.
)

const (
	UnitPercentage              = "percentage" // %
	UnitArcDegrees              = "arcdegrees" // °
	UnitCelsius                 = "celsius"    // °C
	UnitLux                     = "lux"        // lux
	UnitSeconds                 = "seconds"    // sec
	UnitPPM                     = "ppm"        // ppm
	UnitMicrogramsPerCubicMeter = "micrograms/m^3"
)

const (
	FormatString = "string"
	FormatBool   = "bool"
	FormatFloat  = "float"
	FormatUInt8  = "uint8"
	FormatUInt16 = "uint16"
	FormatUInt32 = "uint32"
	FormatInt32  = "int32"
	FormatUInt64 = "uint64"
	FormatData   = "data"
	FormatTLV8   = "tlv8"
)

// ValueUpdateFunc is the value updated function for a characteristic.
type ValueUpdateFunc func(c *C, new, old interface{}, req *http.Request)

// C is a characteristic
type C struct {
	// Id is the unique identifier
	Id uint64

	// Type is the characteristic type (ex. "8" for brightness)
	Type string

	// Permissions are the permissions
	Permissions []string

	// Description is a custom description
	Description string

	// Val is the stored value
	Val interface{}

	// Format is the value format (FormatString, FormatBool, ...)
	Format string

	// Unit is the value unit (UnitPercentage, UnitArcDegrees, ...)
	Unit string

	// MaxLen is the maximum length of Val (maximum characters if the format is "string")
	MaxLen int

	// MaxVal is the maximum value of Val (only for integers and floats)
	MaxVal interface{}

	// MinVal is the minimum value of Val (only for integers and floats)
	MinVal interface{}

	// StepVal is the step value of Val (only for integers and floats)
	StepVal interface{}

	// ValidVals are the valid values for integer characteristics.
	ValidVals []int

	// ValidRange is a 2 element array the valid range start and end.
	ValidRange []int

	// ValueRequestFunc is called when the value of C is requested by a
	// paired controller via an HTTP request.
	// If the value of C represents the state of a remote object, you can use
	// this function to communicate with that object (ex. over the network).
	// If the communication fails, you can return a code != 0.
	// In this case, the server responds with the HTTP status code 500 and the code
	// in the response body (as defined in HAP-R2 6.7.1.4 HAP Status Codes).
	ValueRequestFunc func(request *http.Request) (value interface{}, code int)

	// SetValueRequestFunc is called when the value of C is updated by an
	// HTTP request coming from a paired controller.
	// If the value of C represents the state of a remote object, you can use
	// this function to communicate with that object (ex. over the network).
	// If the communication fails, you can return a code != 0.
	// In this case, the server responds with the HTTP status code 500 and the code
	// in the response body (as defined in HAP-R2 6.7.1.4 HAP Status Codes).
	SetValueRequestFunc func(value interface{}, request *http.Request) (response interface{}, code int)

	// A list of update value functions.
	// There are called when the value of the characteristic is updated.
	valUpdateFuncs []ValueUpdateFunc

	// Flag indicating if the value should be updated even
	// when the new value is the same as the old value.
	// This flag is only used for programmable switch events.
	updateOnSameValue bool

	// Stores which connected client has events enabled for this characteristic.
	events map[string]bool

	m sync.Mutex
}

// New returns a new characteristic.
func New() *C {
	return &C{
		events:         make(map[string]bool),
		valUpdateFuncs: make([]ValueUpdateFunc, 0),
	}
}

// OnCValueUpdate register the given function which is called
// when the value of the characteristic is updated.
func (c *C) OnCValueUpdate(fn ValueUpdateFunc) {
	c.m.Lock()
	c.valUpdateFuncs = append(c.valUpdateFuncs, fn)
	c.m.Unlock()
}

// Sets the value of c to val and returns a status code.
// The server invokes this function when the value is updated by an http request.
func (c *C) SetValueRequest(val interface{}, req *http.Request) (interface{}, int) {
	// check write permission
	if req != nil && !c.IsWritable() {
		log.Info.Printf("writing %v by %s not allowed\n", val, req.RemoteAddr)
		return val, -70404
	}

	return c.setValue(val, req)
}

func (c *C) setValue(v interface{}, req *http.Request) (interface{}, int) {
	newVal := c.convert(v)
	response := newVal
	// Value must be within min and max
	switch c.Format {
	case FormatFloat:
		newVal = c.clampFloat(newVal.(float64))
	case FormatUInt8, FormatUInt16, FormatUInt32, FormatUInt64, FormatInt32:
		newVal = c.clampInt(newVal.(int))
	}

	c.m.Lock()
	// reference old value
	oldVal := c.Val
	c.m.Unlock()

	// ignore the same newVal
	if oldVal == newVal && !c.updateOnSameValue {
		// no error
		return nil, 0
	}

	if !c.validVal(newVal) {
		return nil, -70410
	}

	if c.SetValueRequestFunc != nil && req != nil {
		v, c := c.SetValueRequestFunc(newVal, req)
		if c != 0 {
			return v, c
		}

		if v != nil {
			response = v
		}
	}

	c.m.Lock()
	// update to new value
	c.Val = newVal
	funcs := c.valUpdateFuncs
	c.m.Unlock()

	// call update funcs
	for _, fn := range funcs {
		fn(c, newVal, oldVal, req)
	}

	return response, 0
}

// ValueRequest returns the value of C and a status code.
// If the value of c cannot be read (because it is writeonly),
// the status code -70405 is returned.
func (c *C) ValueRequest(req *http.Request) (interface{}, int) {
	// check write permission
	if !c.IsReadable() {
		log.Info.Printf("reading %d by %s not allowed\n", c.Id, req.RemoteAddr)
		return nil, -70405
	}

	if c.ValueRequestFunc != nil {
		return c.ValueRequestFunc(req)
	}

	return c.Value(), 0
}

// Value returns the value of C
func (c *C) Value() interface{} {
	c.m.Lock()
	defer c.m.Unlock()
	return c.Val
}

func (c *C) SetEvent(remoteAddr string, enable bool) {
	c.m.Lock()
	defer c.m.Unlock()
	c.events[remoteAddr] = enable
}

func (c *C) HasEventsEnabled(remoteAddr string) bool {
	c.m.Lock()
	defer c.m.Unlock()

	ev, ok := c.events[remoteAddr]
	if ok {
		return ev
	}
	return false
}

// IsWritable returns true if clients are allowed
// to update the value of the characteristic.
func (c *C) IsWritable() bool {
	for _, p := range c.Permissions {
		if p == PermissionWrite {
			return true
		}
	}

	return false
}

// IsReadable returns true if clients are allowed
// to read the value of the characteristic.
func (c *C) IsReadable() bool {
	for _, p := range c.Permissions {
		if p == PermissionRead {
			return true
		}
	}

	return false
}

// RequiresTimedWrite returns true if the value can
// only be set with a timed write procedure.
func (c *C) RequiresTimedWrite() bool {
	for _, p := range c.Permissions {
		if p == PermissionTimedWrite {
			return true
		}
	}

	return false
}

// IsWriteResponse returns true if the value can
// return a response on write
func (c *C) IsWriteResponse() bool {
	for _, p := range c.Permissions {
		if p == PermissionWriteResponse {
			return true
		}
	}

	return false
}

// IsObservable returns true if clients are allowed
// to observe the value of the characteristic.
func (c *C) IsObservable() bool {
	for _, p := range c.Permissions {
		if p == PermissionEvents {
			return true
		}
	}

	return false
}

// IsObservable returns true if the value of the
// characteristic can only be updated, but not read.
func (c *C) IsWriteOnly() bool {
	return len(c.Permissions) == 1 && c.Permissions[0] == PermissionWrite
}

func (c *C) MarshalJSON() ([]byte, error) {
	d := struct {
		Id          uint64   `json:"iid"` // managed by accessory
		Type        string   `json:"type"`
		Permissions []string `json:"perms"`
		Format      string   `json:"format"`

		Value       *V          `json:"value,omitempty"`
		Description string      `json:"description,omitempty"` // manufacturer description (optional)
		Unit        string      `json:"unit,omitempty"`
		MaxLen      int         `json:"maxLen,omitempty"`
		MaxValue    interface{} `json:"maxValue,omitempty"`
		MinValue    interface{} `json:"minValue,omitempty"`
		StepValue   interface{} `json:"minStep,omitempty"`
		ValidValues []int       `json:"valid-values,omitempty"`
		ValidRange  []int       `json:"valid-values-range,omitempty"`
	}{
		Id:          c.Id,
		Type:        c.Type,
		Permissions: c.Permissions,
		Description: c.Description,
		Format:      c.Format,
		Unit:        c.Unit,
		MaxLen:      c.MaxLen,
		MaxValue:    c.MaxVal,
		MinValue:    c.MinVal,
		StepValue:   c.StepVal,
		ValidValues: c.ValidVals,
		ValidRange:  c.ValidRange,
	}

	// If the characteristic is readable, the value
	// must be present in the json representation.
	if c.IsReadable() {
		// 2022-03-21 (mah) FIXME provide a http request instead of nil
		if v, s := c.ValueRequest(nil); s == 0 {
			d.Value = &V{v}
		} else {
			d.Value = &V{c.Value()} // dummy "zero" value
		}
	}

	return json.Marshal(&d)
}

type V struct {
	Value interface{}
}

func (v V) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Value)
}

func (c *C) clampFloat(value float64) interface{} {
	min, minOK := c.MinVal.(float64)
	max, maxOK := c.MaxVal.(float64)
	if maxOK == true && value > max {
		value = max
	} else if minOK == true && value < min {
		value = min
	}

	return value
}

func (c *C) clampInt(value int) interface{} {
	min, minOK := c.MinVal.(int)
	max, maxOK := c.MaxVal.(int)
	if maxOK == true && value > max {
		value = max
	} else if minOK == true && value < min {
		value = min
	}

	return value
}

func (c *C) convert(v interface{}) interface{} {
	switch c.Format {
	case FormatFloat:
		return to.Float64(v)
	case FormatUInt8, FormatUInt16, FormatUInt32, FormatInt32:
		return int(to.Uint64(v))
	case FormatUInt64:
		return to.Uint64(v)
	case FormatBool:
		return to.Bool(v)
	default:
		return v
	}
}

func (c *C) validVal(v interface{}) bool {
	iv, ok := v.(int)
	if !ok {
		return true
	}

	if len(c.ValidVals) > 0 {
		for _, val := range c.ValidVals {
			if val == v {
				return true
			}
		}

		return false
	}

	if len(c.ValidRange) == 2 {
		return c.ValidRange[0] <= iv && c.ValidRange[1] >= iv
	}

	return true
}
