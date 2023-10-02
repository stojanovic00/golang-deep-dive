package main

import "testing"

func TestCreateFormSuccess(t *testing.T) {
	data := []string{"David", "9.20", "Alex", "8.10", "Max", "6.20", "Ben", "7.50"}

	got, err := CreateForm(data...)
	if err != nil {
		t.Error(err)
	}

	want := map[string]float64{"David": 9.20, "Alex": 8.10, "Max": 6.20, "Ben": 7.50}

	for key, value := range got {
		wanted, ok := want[key]
		if !ok {
			t.Errorf("Key %q is not present in map", key)
			return
		}
		if value != wanted {
			t.Errorf("Expected %v for key %q, but got %v", wanted, key, value)
			return
		}
	}
}

func TestCreateFormOddParameters(t *testing.T) {
	data := []string{"David", "9.20", "Alex", "8.10", "Max", "6.20", "Ben"}

	_, err := CreateForm(data...)

	if err != ErrOddParameterNumber {
		t.Errorf("Expected \"ErrOddParameterNumber\", got: %v", err)
	}
}

func TestCreateFormNotANumber(t *testing.T) {
	data := []string{"David", "9.20", "Alex", "Steven", "Max", "6.20", "Ben", "7.50"}

	_, err := CreateForm(data...)

	if _, ok := err.(ErrNotANumber); !ok {
		t.Errorf("Expected \"ErrNotANumber\", got %v", err)
	}
}
