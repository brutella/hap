package characteristic

import (
	"github.com/brutella/hap/log"
	"github.com/xiam/to"

	"encoding/json"
	"net/http"
)

const (
	PermissionRead          = "pr" // The characteristic can only be read by paired controllers.
	PermissionWrite         = "pw" // The characteristic can only be written by paired controllers.
	PermissionEvents        = "ev" // The characteristic supports events.
	PermissionHidden        = "hd" // The characteristic is hidden from the user
	PermissionWriteResponse = "wr" // The characteristic supports write response
)

const (
	UnitPercentage = "percentage" // %
	UnitArcDegrees = "arcdegrees" // °
	UnitCelsius    = "celsius"    // °C
	UnitLux        = "lux"        // lux
	UnitSeconds    = "seconds"    // sec
	UnitPPM        = "ppm"        // ppm
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

	// Stores which connected client has events enabled for this characteristic.
	Events map[string]bool

	// ValueRequestFunc is called when the value of C is requested.
	// The http request is non-nil, if the value of C is requested.
	// by an HTTP request coming from a paired controller.
	// The first return value of this function is the value of C.
	// If the second argument is non-zero, the server responds with the
	// HTTP status code 500 Internal Server Error.
	ValueRequestFunc func(request *http.Request) (interface{}, int)

	// SetValueRequestFunc is called when the value of C is updated.
	// The first argument "value" is the new value of C.
	// The second argument "request" is non-nil, if the value of C is
	// updated from an HTTP request coming from a paired controller.
	// An error is inidcated if the return value is non-zero.
	SetValueRequestFunc func(value interface{}, request *http.Request) int

	// A list of update value functions.
	// There are called when the value of the characteristic is updated.
	valUpdateFuncs []ValueUpdateFunc

	// Flag indicating if the value should be updated even
	// when the new value is the same as the old value.
	// This flag is only used for programmable switch events.
	updateOnSameValue bool
}

// New returns a new characteristic.
func New() *C {
	return &C{
		Events:         make(map[string]bool),
		valUpdateFuncs: make([]ValueUpdateFunc, 0),
	}
}

// OnCValueUpdate register the provided function, which is called
// when the value of the characteristic is updated.
func (c *C) OnCValueUpdate(fn ValueUpdateFunc) {
	c.valUpdateFuncs = append(c.valUpdateFuncs, fn)
}

// Sets the value of c to v.
// The function is called if the value is updated from an http request.
func (c *C) SetValueRequest(v interface{}, req *http.Request) int {
	// check write permission
	if !c.IsWritable() {
		log.Info.Printf("writing %v by %s not allowed\n", v, req.RemoteAddr)
		return -70404
	}

	if c.SetValueRequestFunc != nil {
		if s := c.SetValueRequestFunc(v, req); s != 0 {
			return s
		}
	}

	c.setValue(v, req)

	// no error
	return 0
}

func (c *C) setValue(v interface{}, req *http.Request) {
	newVal := c.convert(v)

	// Value must be within min and max
	switch c.Format {
	case FormatFloat:
		newVal = c.clampFloat(newVal.(float64))
	case FormatUInt8, FormatUInt16, FormatUInt32, FormatUInt64, FormatInt32:
		newVal = c.clampInt(newVal.(int))
	}

	// ignore the same newVal
	if c.Val == newVal && !c.updateOnSameValue {
		return
	}

	// reference old value
	oldVal := c.Val

	// update to new value
	c.Val = newVal

	// call update funcs
	for _, fn := range c.valUpdateFuncs {
		fn(c, newVal, oldVal, req)
	}
}

// ValueRequest returns the value of C and a status code.
// if the characteristic is not readable, the status code -70405 is returned.
func (c *C) ValueRequest(req *http.Request) (interface{}, int) {
	// check write permission
	if !c.IsReadable() {
		log.Info.Printf("reading %d by %s not allowed\n", c.Id, req.RemoteAddr)
		return nil, -70405
	}

	return c.valueRequest(req)
}

// value returns the value of C and a status code.
func (c *C) valueRequest(req *http.Request) (interface{}, int) {
	if c.ValueRequestFunc != nil {
		return c.ValueRequestFunc(req)
	}

	return c.Val, 0
}

func (c *C) IsWritable() bool {
	for _, p := range c.Permissions {
		if p == PermissionWrite {
			return true
		}
	}

	return false
}

func (c *C) IsReadable() bool {
	for _, p := range c.Permissions {
		if p == PermissionRead {
			return true
		}
	}

	return false
}

func (c *C) IsObservable() bool {
	for _, p := range c.Permissions {
		if p == PermissionEvents {
			return true
		}
	}

	return false
}

func (c *C) IsWriteOnly() bool {
	return len(c.Permissions) == 1 && c.Permissions[0] == PermissionWrite
}

func (c *C) MarshalJSON() ([]byte, error) {
	d := struct {
		Id          uint64   `json:"iid"` // managed by accessory
		Type        string   `json:"type"`
		Permissions []string `json:"perms"`
		Description string   `json:"description,omitempty"` // manufacturer description (optional)

		Value  interface{} `json:"value,omitempty"` // nil for write-only characteristics
		Format string      `json:"format"`
		Unit   string      `json:"unit,omitempty"`

		MaxLen    int         `json:"maxLen,omitempty"`
		MaxValue  interface{} `json:"maxValue,omitempty"`
		MinValue  interface{} `json:"minValue,omitempty"`
		StepValue interface{} `json:"minStep,omitempty"`
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
	}

	if c.IsReadable() {
		if v, _ := c.valueRequest(nil); v != nil {
			d.Value = v
		}
	}

	return json.Marshal(&d)
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
