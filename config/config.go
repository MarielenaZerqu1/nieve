package config

import(
	"encoding/json"
	"path/filepath"
	"bytes"
	"io/ioutil"

	"github.com/dilmorja/nieve/spdx"
	"github.com/dilmorja/nieve/internal/errors"
)

// The Layout are the configuration options.
// From here you can access the values of these options.
type Layout struct {
	Name string `json:"name"`
	Version string `json:"version"`
	License string `json:"license"`
}

// Load options from a JSON file.
func (lay *Layout) Load(file string) error {
	absFilePath, err := filepath.Abs(file)

	if err != nil {
		return err
	}

	var fileContent []byte
	
	fileContent, err = ioutil.ReadFile(absFilePath)

	if err != nil {
		return err
	}

	var decoder = json.NewDecoder(bytes.NewReader(fileContent))
	decoder.DisallowUnknownFields()

	if err = decoder.Decode(lay); err != nil {
		return err
	}

	var idOk bool
	if idOk, err = spdx.Validate(lay.License); !idOk && err == nil {
		return errors.ErrInvalidSPDXLicenseId
	} else {
		return err
	}

	return nil
}