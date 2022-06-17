package config

import(
	"testing"
)

func TestLayoutLoadFrom(t *testing.T) {
	expected := []interface{}{
		"Test",
		"0.0.0",
		"BSD-3-Clause",
	}

	var data Layout

	err := data.LoadFrom("../testdata/testLayout.json")

	if err != nil {
		t.Fatalf("config.Layout.Load(): %v", err)
	}

	got := []interface{}{
		data.Name,
		data.Version,
		data.License,
	}

	for i, e := range expected {
		if e != got[i] {
			t.Errorf("\nExpected (%d): %v\nGot: %v\n", i, e, got[i])
		}
	}
}