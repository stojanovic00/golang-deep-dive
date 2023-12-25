package main

import (
	"reflect"
	"testing"
)

func TestSortAndGroup(t *testing.T) {
	numbers := []int{20, 40, 10, 50, 80, 60, 90, 70, 30, 100}

	got := SortAndGroup(numbers)

	want := [][]int{{10, 20, 30}, {40, 50, 60}, {70, 80, 90, 100}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected: %v, got: %v ", want, got)
	}
}
