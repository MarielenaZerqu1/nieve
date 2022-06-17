// SPDX-License-Identifier: BSD-4-Clause

package errors

import "errors"

var(
	ErrInvalidSPDXLicenseId = errors.New("License identifier is not an SPDX standard. (https://spdx.org/licenses)")
)