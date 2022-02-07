// +build ignore

// Imports HomeKit metadata from HomeKitDaemon.framework
package main

import (
	"encoding/json"
	"fmt"
	"github.com/brutella/hap/gen"
	_ "github.com/brutella/hap/gen/golang"
	"github.com/brutella/hap/gen/meta"
	"io/ioutil"
	"log"
	"os"
	_ "os/exec"
	"path/filepath"
)

var LibPath = os.ExpandEnv("$GOPATH/src/github.com/brutella/hap")
var GenPath = filepath.Join(LibPath, "gen")
var SvcPkgPath = filepath.Join(LibPath, "service")
var AccPkgPath = filepath.Join(LibPath, "accessory")
var CharPkgPath = filepath.Join(LibPath, "characteristic")
var MetadataPath = filepath.Join(GenPath, "plain-metadata.json")

func findChar(id string, data meta.Data) *meta.Char {
	for key, ch := range data.Hap.Chars {
		if key == id {
			return &ch
		}
	}

	return nil
}

func validValues(id string, data meta.Data) *map[string]string {
	for _, ch := range data.Assistant.Chars {
		if ch.Read == id || ch.Write == id || ch.ReadWrite == id {
			val := map[string]string{}
			for key, value := range ch.Values {
				newKey := fmt.Sprintf("%v", value)
				val[newKey] = key
			}

			return &val
		}
	}

	return nil
}

func main() {

	log.Println("Import data from", MetadataPath)

	// Open metadata file
	f, err := os.Open(MetadataPath)
	if err != nil {
		log.Fatal(err)
	}

	// Read content
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	// Import json
	data := meta.Data{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		log.Fatal(err)
	}

	chars := []*gen.CharacteristicMetadata{}
	// Create characteristic files
	for id, char := range data.Hap.Chars {
		constraints := map[string]interface{}{}
		if char.Step != nil {
			constraints["StepValue"] = *char.Step
		}

		if char.Max != nil {
			constraints["MaximumValue"] = *char.Max
		}

		if char.Min != nil {
			constraints["MinimumValue"] = *char.Min
		}

		if val := validValues(id, data); val != nil && len(*val) > 0 {
			constraints["ValidValues"] = val
		}

		props := []string{}
		switch char.Properties {
		case 3:
			props = []string{"read", "cnotify"}
		case 4:
			props = []string{"write"}
		case 7:
			props = []string{"read", "write", "cnotify"}
		default:
			break
		}

		unit := ""
		if char.Unit != nil {
			unit = *char.Unit
		}

		c := gen.CharacteristicMetadata{
			Constraints: constraints,
			Format:      char.Format,
			Name:        char.Description,
			UUID:        fmt.Sprintf("%s%s", char.UUID, data.Hap.BaseUUID),
			Unit:        unit,
		}
		chars = append(chars, &c)
	}

	svs := []*gen.ServiceMetadata{}
	// Create characteristic files
	for _, sv := range data.Hap.Services {
		required := []string{}
		for _, id := range sv.Char.Required {
			ch := findChar(id, data)
			uuid := fmt.Sprintf("%s%s", ch.UUID, data.Hap.BaseUUID)
			required = append(required, uuid)
		}
		optional := []string{}
		for _, id := range sv.Char.Optional {
			ch := findChar(id, data)
			if ch == nil {
				log.Println("Unable to find %s", ch.Description)
				continue
			}
			uuid := fmt.Sprintf("%s%s", ch.UUID, data.Hap.BaseUUID)
			optional = append(optional, uuid)
		}

		uuid := fmt.Sprintf("%s%s", sv.UUID, data.Hap.BaseUUID)
		sv := gen.ServiceMetadata{
			RequiredCharacteristics: required,
			OptionalCharacteristics: optional,
			Name:                    sv.Description,
			UUID:                    uuid,
		}
		svs = append(svs, &sv)
	}

	metadata := gen.Metadata{
		nil, chars, svs,
	}
	b, err = json.Marshal(metadata)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
