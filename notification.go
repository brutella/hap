package hap

import (
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/characteristic"
	"github.com/brutella/hap/log"

	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func sendNotification(a *accessory.A, c *characteristic.C, req *http.Request) error {
	pl := struct {
		Cs []characteristicData `json:"characteristics"`
	}{
		Cs: []characteristicData{
			characteristicData{
				Aid:   a.Id,
				Iid:   c.Id,
				Value: &characteristic.V{c.Val},
			},
		},
	}

	plb, err := json.Marshal(pl)
	if err != nil {
		return err
	}

	body := bytes.NewBuffer(plb)

	// Use http.Response to send the notification as an http message.
	resp := new(http.Response)
	resp.Status = "200 OK"
	resp.StatusCode = http.StatusOK
	resp.ProtoMajor = 1
	resp.ProtoMinor = 0
	resp.Body = ioutil.NopCloser(body)
	resp.ContentLength = int64(body.Len())
	resp.Header = map[string][]string{}
	resp.Header.Set("Content-Type", HTTPContentTypeHAPJson)

	// Will be ignored unfortunately and won't be fixed https://github.com/golang/go/issues/9304
	// Make sure to call FixProtocolSpecifier() instead
	// resp.Proto = "EVENT/1.0"

	// Set protocol of message to "EVENT/1.0".
	var buffer = new(bytes.Buffer)
	resp.Write(buffer)
	b, err := ioutil.ReadAll(buffer)
	b = []byte(strings.Replace(string(b), "HTTP/1.0", "EVENT/1.0", 1))

	for _, conn := range conns() {
		if req != nil && req.RemoteAddr == conn.RemoteAddr().String() {
			// Don't send notification to the client
			// who updated the value.
			log.Debug.Printf("skip notification for %s\n", conn.RemoteAddr())
			continue
		}

		// Check which connection has events enabled.
		if c.HasEventsEnabled(conn.RemoteAddr().String()) {
			log.Debug.Printf("send event to %s:\n%s\n", conn.RemoteAddr(), string(b))
			conn.Write(b)
		}
	}

	return nil
}
