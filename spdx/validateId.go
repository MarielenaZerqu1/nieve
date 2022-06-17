// SPDX-License-Identifier: BSD-4-Clause

package spdx

import(
	"net/http"
	"io/ioutil"
	"encoding/json"
)

const spdxJSON_URL string = "https://spdx.org/licenses/licenses.json"

func Validate(targetID string) (bool, error) {
	// Make request to spdxJSON_URL
	res, err := http.Get(spdxJSON_URL)

	if err != nil {
		return false, err
	}

	// Parse the response body
	var licensesBuf []byte
	licensesBuf, err = ioutil.ReadAll(res.Body)

	if err != nil {
		return false, err
	}

	// Unmarshal json
	var data LicensesList

	err = json.Unmarshal(licensesBuf, &data)

	if err != nil {
		return false, err
	}

	// Verify
	for _, license := range data.Licenses {
		if license.ID == targetID {
			return true, nil
		}
	}

	return false, nil
}