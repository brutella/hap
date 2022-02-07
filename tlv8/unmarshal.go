package tlv8

import (
	"io"
	"io/ioutil"
)

func UnmarshalReader(r io.Reader, v interface{}) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	return unmarshal(b, v)
}

func Unmarshal(data []byte, v interface{}) error {
	return unmarshal(data, v)
}

func unmarshal(data []byte, v interface{}) error {
	d, err := newDecoder(data)
	if err != nil {
		return err
	}

	return d.decode(v)
}
