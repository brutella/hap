package characteristic

import (
	"github.com/brutella/hap/log"
	"github.com/xiam/to"

	"encoding/json"
	"net/http"
)

const (
	PermissionRead          = "pr" // can be read
	PermissionWrite         = "pw" // can be written
	PermissionEvents        = "ev" // sends events
	PermissionHidden        = "hd" // is hidden
	PermissionWriteResponse = "wr"
)

const (
	UnitPercentage = "percentage"
	UnitArcDegrees = "arcdegrees"
	UnitCelsius    = "celsius"
	UnitLux        = "lux"
	UnitSeconds    = "seconds"
	UnitPPM        = "ppm"
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

type ValueUpdateFunc func(c *C, new, old interface{}, req *http.Request)

// C is a characteristic
type C struct {
	Id          uint64
	Type        string
	Permissions []string
	Description string
	Val         interface{}
	Format      string
	Unit        string
	MaxLen      int
	MaxVal      interface{}
	MinVal      interface{}
	StepVal     interface{}

	// Maps events flags to address
	Events map[string]bool

	// ValFunc returns the value of C.
	// If no nil, the return value of this function is used instead of Val.
	// req is non nil, if the value is requested from a request.
	ValFunc func(req *http.Request) interface{}

	// A list of update value functions.
	// There are called when the value of the characteristic is updated.
	valUpdateFuncs []ValueUpdateFunc

	updateOnSameValue bool
}

func New() *C {
	return &C{
		Events:         make(map[string]bool),
		valUpdateFuncs: make([]ValueUpdateFunc, 0),
	}
}

func (c *C) OnCValueUpdate(fn ValueUpdateFunc) {
	c.valUpdateFuncs = append(c.valUpdateFuncs, fn)
}

// value returns the value of the characteristic.
func (c *C) value(r *http.Request) interface{} {
	if c.ValFunc != nil {
		return c.ValFunc(r)
	}

	return c.Val
}

// Sets the value of c to v.
// The function is called if the value is updated from an http request.
func (c *C) SetValueRequest(v interface{}, req *http.Request) {
	// check write permission
	if !c.IsWritable() {
		log.Info.Printf("writing %v by %s not allowed\n", v, req.RemoteAddr)
		return
	}
	c.setValue(v, req)
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

func (c *C) ValueRequest(req *http.Request) interface{} {
	// check write permission
	if !c.IsReadable() {
		log.Info.Printf("reading %d by %s not allowed\n", c.Id, req.RemoteAddr)
		return nil
	}

	if c.ValFunc != nil {
		return c.ValFunc(req)
	}

	return c.Val
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

func (ch *C) MarshalJSON() ([]byte, error) {
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
		Id:          ch.Id,
		Type:        ch.Type,
		Permissions: ch.Permissions,
		Description: ch.Description,
		Format:      ch.Format,
		Unit:        ch.Unit,
		MaxLen:      ch.MaxLen,
		MaxValue:    ch.MaxVal,
		MinValue:    ch.MinVal,
		StepValue:   ch.StepVal,
	}

	if ch.IsReadable() {
		d.Value = ch.value(nil)
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
