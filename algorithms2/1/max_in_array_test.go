package main

import (
	"reflect"
	"testing"
)

func TestMaxInArray(t *testing.T) {
	arrays := [][]int{{32, 12, 24, 20}, {18, 40, 22, 30}, {21, 31, 42, 35}}

	got := MaxInArray(arrays)

	want := []int{32, 40, 42}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected: %v, got: %v", want, got)
	}
}
