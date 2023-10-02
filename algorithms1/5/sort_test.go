package main

import (
	"reflect"
	"testing"
)

// TODO sometimes it passes, sometimes not.
// Identical includes also having elements in the same order
func IdenticalMaps(m1, m2 map[string]int) bool {

	//Checking key:value equality
	m1Keys := make([]string, 0, len(m1))
	for key, value := range m1 {
		m1Keys = append(m1Keys, key)
		//Doesn't even have all the keys
		if _, ok := m2[key]; !ok {
			return false
		}

		if m2[key] != value {
			return false
		}
	}

	//Checking element order
	m2Keys := make([]string, 0, len(m2))
	for key, _ := range m2 {
		m2Keys = append(m2Keys, key)
	}

	return reflect.DeepEqual(m1Keys, m2Keys)
}

func TestSortByKey(t *testing.T) {
	data := map[string]int{"David": 40, "Paul": 20, "Bill": 30, "Fred": 50, "Alex": 10}

	got := SortByKey(data)

	want := map[string]int{"Alex": 10, "Bill": 30, "David": 40, "Fred": 50, "Paul": 20}

	if !IdenticalMaps(got, want) {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}

func TestSortByValue(t *testing.T) {
	data := map[string]int{"David": 40, "Paul": 20, "Bill": 30, "Fred": 50, "Alex": 10}

	got := SortByValue(data)

	want := map[string]int{"Alex": 10, "Paul": 20, "Bill": 30, "David": 40, "Fred": 50}

	if !IdenticalMaps(got, want) {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}
