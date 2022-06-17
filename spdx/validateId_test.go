// SPDX-License-Identifier: BSD-4-Clause

package spdx

import "testing"

func TestValidate(t *testing.T) {
	expected := true

	got, err := Validate("MIT")

	if err != nil {
		t.Fatalf("spdx.Validate(): %v", err)
	}

	if expected != got {
		t.Errorf("\nExpected: %v\nGot: %v\n", expected, got)
	}
}