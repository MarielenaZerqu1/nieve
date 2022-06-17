// SPDX-License-Identifier: BSD-4-Clause

package spdx

// Use only the license ID
type License struct {
	ID string `json:"licenseId"`
}

// The main JSON. Use only the list of licenses.
type LicensesList struct {
	Licenses []License `json:"licenses"`
}