package hap

import (
	"time"

	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/characteristic"
	"github.com/brutella/hap/log"
	"github.com/xiam/to"

	"encoding/json"
	"net/http"
	"strings"
)

type characteristicData struct {
	Aid   uint64            `json:"aid"`
	Iid   uint64            `json:"iid"`
	Value *characteristic.V `json:"value,omitempty"`

	// optional values
	Type        *string     `json:"type,omitempty"`
	Permissions []string    `json:"perms,omitempty"`
	Status      *int        `json:"status,omitempty"`
	Events      *bool       `json:"ev,omitempty"`
	Format      *string     `json:"format,omitempty"`
	Unit        *string     `json:"unit,omitempty"`
	MinValue    interface{} `json:"minValue,omitempty"`
	MaxValue    interface{} `json:"maxValue,omitempty"`
	MinStep     interface{} `json:"minStep,omitempty"`
	MaxLen      *int        `json:"maxLen,omitempty"`
	ValidValues []int       `json:"valid-values,omitempty"`
	ValidRange  []int       `json:"valid-values-range,omitempty"`
}

type putCharacteristicData struct {
	Aid uint64 `json:"aid"`
	Iid uint64 `json:"iid"`

	Value  interface{} `json:"value,omitempty"`
	Status *int        `json:"status,omitempty"`
	Events *bool       `json:"ev,omitempty"`

	Remote   *bool `json:"remote,omitempty"`
	Response *bool `json:"r,omitempty"`
}

func (srv *Server) getCharacteristics(res http.ResponseWriter, req *http.Request) {
	if !srv.IsAuthorized(req) {
		log.Info.Printf("request from %s not authorized\n", req.RemoteAddr)
		JsonError(res, JsonStatusInsufficientPrivileges)
		return
	}

	// id=1.4,1.5
	v := req.FormValue("id")
	if len(v) == 0 {
		JsonError(res, JsonStatusInvalidValueInRequest)
		return
	}

	meta := req.FormValue("meta") == "1"
	perms := req.FormValue("perms") == "1"
	typ := req.FormValue("type") == "1"
	ev := req.FormValue("ev") == "1"

	arr := []*characteristicData{}
	err := false
	for _, str := range strings.Split(v, ",") {
		ids := strings.Split(str, ".")
		if len(ids) != 2 {
			continue
		}
		cdata := &characteristicData{
			Aid: to.Uint64(ids[0]),
			Iid: to.Uint64(ids[1]),
		}
		arr = append(arr, cdata)

		c := srv.findC(cdata.Aid, cdata.Iid)
		if c == nil {
			err = true
			status := JsonStatusServiceCommunicationFailure
			cdata.Status = &status
			continue
		}

		v, s := c.ValueRequest(req)
		if s != 0 {
			err = true
			cdata.Status = &s
		} else {
			cdata.Value = &characteristic.V{v}
		}

		if meta {
			cdata.Format = &c.Format
			cdata.Unit = &c.Unit
			if c.MinVal != nil {
				cdata.MinValue = c.MinVal
			}
			if c.MaxVal != nil {
				cdata.MaxValue = c.MaxVal
			}
			if c.StepVal != nil {
				cdata.MinStep = c.StepVal
			}

			if c.MaxLen > 0 {
				cdata.MaxLen = &c.MaxLen
			}

			if len(c.ValidVals) > 0 {
				cdata.ValidValues = c.ValidVals
			}

			if len(c.ValidRange) > 0 {
				cdata.ValidRange = c.ValidRange
			}
		}

		// Should the response include the events flag?
		if ev {
			ev := c.HasEventsEnabled(req.RemoteAddr)
			cdata.Events = &ev
		}

		if perms {
			cdata.Permissions = c.Permissions
		}

		if typ {
			cdata.Type = &c.Type
		}
	}

	resp := struct {
		Characteristics []*characteristicData `json:"characteristics"`
	}{arr}

	log.Debug.Println(toJSON(resp))

	if err {
		// when there's an error somewhere, "status: 0" must now be explicit
		noError := 0
		for _, c := range arr {
			if c.Status == nil {
				c.Status = &noError
			}
		}

		JsonMultiStatus(res, resp)
	} else {
		JsonOK(res, resp)
	}
}

func (srv *Server) putCharacteristics(res http.ResponseWriter, req *http.Request) {
	if !srv.IsAuthorized(req) {
		log.Info.Printf("request from %s not authorized\n", req.RemoteAddr)
		JsonError(res, JsonStatusInsufficientPrivileges)
		return
	}

	data := struct {
		Cs  []putCharacteristicData `json:"characteristics"`
		Pid uint64                  `json:"pid"`
	}{}

	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		JsonError(res, JsonStatusInvalidValueInRequest)
		return
	}

	timedWr := srv.TimedWrite(req)
	log.Debug.Println(toJSON(data))

	arr := []*putCharacteristicData{}
	for _, d := range data.Cs {
		c := srv.findC(d.Aid, d.Iid)

		cdata := &putCharacteristicData{
			Aid: d.Aid,
			Iid: d.Iid,
		}

		if c == nil {
			status := JsonStatusServiceCommunicationFailure
			cdata.Status = &status
			arr = append(arr, cdata)
			continue
		}

		var value interface{}
		var status int
		if c.RequiresTimedWrite() {
			if time.Now().After(timedWr.deadline) {
				// HAP 6.7.2.4
				// If the accessory receives an Execute Write Request after the TTL has expired it must ignore
				// the request and respond with HAP status error code -70410 (HAPIPStatusErrorCodeInvalidWrite).
				log.Info.Println("timed write wall time exceeded")
				status = -70410
			}
			if data.Pid != timedWr.pid {
				// HAP 6.7.2.4
				// If the accessory receives a standard write request on a characteristic which requires timed write,
				// the accessory must respond with HAP status error code -70410 (HAPIPStatusErrorCodeInvalidWrite).
				log.Info.Println("timed write transaction id invalid")
				status = -70410
			}
		}

		if d.Value != nil && status == 0 {
			value, status = c.SetValueRequest(d.Value, req)
		}

		if status != 0 {
			cdata.Status = &status
		}

		if (d.Response != nil || c.IsWriteResponse()) && value != nil {
			cdata.Value = value

			if c.IsWriteResponse() {
				cdata.Status = &status
			}
		}

		if d.Events != nil {
			if !c.IsObservable() {
				status := JsonStatusNotificationNotSupported
				cdata.Status = &status
				arr = append(arr, cdata)
			} else {
				c.SetEvent(req.RemoteAddr, *d.Events)
			}
		}

		if cdata.Status != nil || cdata.Value != nil {
			arr = append(arr, cdata)
		}
	}

	srv.DelTimedWrite(req)

	if len(arr) == 0 {
		res.WriteHeader(http.StatusNoContent)
		return
	}

	resp := struct {
		Characteristics []*putCharacteristicData `json:"characteristics"`
	}{arr}

	log.Debug.Println(toJSON(resp))
	JsonMultiStatus(res, resp)
}

func (srv *Server) findC(aid, iid uint64) *characteristic.C {
	var as []*accessory.A
	as = append(as, srv.a)
	as = append(as, srv.as[:]...)

	for _, a := range as {
		if a.Id == aid {
			for _, s := range a.Ss {
				for _, c := range s.Cs {
					if c.Id == iid {
						return c
					}
				}
			}
		}
	}

	return nil
}

func (srv *Server) prepareCharacteristics(res http.ResponseWriter, req *http.Request) {
	if !srv.IsAuthorized(req) {
		log.Info.Printf("request from %s not authorized\n", req.RemoteAddr)
		JsonError(res, JsonStatusInsufficientPrivileges)
		return
	}

	data := struct {
		Ttl uint64 `json:"ttl"`
		Pid uint64 `json:"pid"`
	}{}

	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil || data.Ttl == 0 || data.Pid == 0 {
		JsonError(res, JsonStatusInvalidValueInRequest)
		return
	}

	srv.SetTimedWrite(data.Ttl, data.Pid, req)

	resp := struct {
		Status int `json:"status"`
	}{0}
	log.Debug.Println(toJSON(resp))
	JsonOK(res, resp)
}
