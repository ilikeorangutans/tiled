package tmx

import (
	"encoding/xml"
	"os"
)

// Loads a Map from the given filename. Returns the map if successfully loaded,
// error if loading or deserialization failed.
func LoadTmxWithFilename(filename string) (*Map, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	decoder := xml.NewDecoder(file)

	m := &Map{}

	err = decoder.Decode(m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
