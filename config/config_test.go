package config

import(
	"testing"
)

func TestLayoutLoadFrom(t *testing.T) {
	expected := []interface{}{
		"Test",
		"A layout test",
		"The Nieve Authors",
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
		data.Description,
		data.Owner,
		data.Version,
		data.License,
	}

	for i, e := range expected {
		if e != got[i] {
			t.Errorf("\nExpected (%d): %v\nGot: %v\n", i, e, got[i])
		}
	}
}

func TestLayoutLoad(t *testing.T) {
	expected := "LT"

	var got Layout

	if err := Load([]byte(`{"name": "LT"}`), &got); err != nil {
		t.Errorf("config.Load(): %v", err)
	}

	if expected != got.Name {
		t.Errorf("\nExpected: %s\nGot: %s\n", expected, got.Name)
	}
}