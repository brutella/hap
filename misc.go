package hap

import (
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/ed25519"
	"github.com/brutella/hap/log"

	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"strings"
)

func configHash(as []*accessory.A) []byte {
	data := struct {
		As []*accessory.A `json:"accessories"`
	}{
		As: as,
	}

	b, err := json.Marshal(data)
	if err != nil {
		log.Info.Panic(err)
	}

	val := map[string]interface{}{}
	if err := json.Unmarshal(b, &val); err != nil {
		log.Info.Panic(err)
	}

	deleteFieldFromDict(&val, "value")

	if b, err = json.Marshal(val); err != nil {
		log.Info.Panic(err)
	}

	h := md5.New()
	h.Write(b)
	return h.Sum(nil)
}

func deleteFieldFromDict(val *map[string]interface{}, field string) {
	for k, v := range *val {
		if k == field {
			delete(*val, k)
		} else {
			deleteFieldFromInterface(&v, field)
		}
	}
}

func deleteFieldFromArray(val *[]interface{}, field string) {
	for _, v := range *val {
		deleteFieldFromInterface(&v, field)
	}
}

func deleteFieldFromInterface(val *interface{}, field string) {
	v := *val

	if dict, ok := v.(map[string]interface{}); ok == true {
		deleteFieldFromDict(&dict, field)
	}

	if array, ok := v.([]interface{}); ok == true {
		deleteFieldFromArray(&array, field)
	}
}

// generateKeyPair generates random public and private key pair
func generateKeyPair() (KeyPair, error) {
	str := randHex()
	public, private, err := ed25519.GenerateKey(str)
	return KeyPair{public[:], private[:]}, err
}

// randomHex returns a random hex string.
func randHex() string {
	var b [16]byte
	// Read might block
	// > crypto/rand: blocked for 60 seconds waiting to read random data from the kernel
	// > https://github.com/golang/go/commit/1961d8d72a53e780effa18bfa8dbe4e4282df0b2
	_, err := rand.Read(b[:])
	if err != nil {
		panic(err)
	}
	var out [32]byte
	for i := 0; i < len(b); i++ {
		out[i*2] = btoh((b[i] >> 4) & 0xF)
		out[i*2+1] = btoh(b[i] & 0xF)
	}
	return string(out[:])
}

func btoh(i byte) byte {
	if i > 9 {
		return 0x61 + (i - 10)
	}
	return 0x30 + i
}

// mac48Address returns a MAC-48-like address from the argument string
func mac48Address(input string) string {
	h := md5.New()
	h.Write([]byte(input))
	result := h.Sum(nil)

	var c []string
	c = append(c, toHex(result[0]))
	c = append(c, toHex(result[1]))
	c = append(c, toHex(result[2]))
	c = append(c, toHex(result[3]))
	c = append(c, toHex(result[4]))
	c = append(c, toHex(result[5]))

	// setup id needs the mac address in upper case
	return strings.ToUpper(strings.Join(c, ":"))
}

func toHex(b byte) string {
	return hex.EncodeToString([]byte{b})
}
