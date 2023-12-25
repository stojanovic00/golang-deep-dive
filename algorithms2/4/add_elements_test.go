package main

import (
	"reflect"
	"testing"
)

func TestAddElements(t *testing.T) {
	list1 := [][]int{{6, 11, 17}, {24, 32, 38}, {45, 50, 55}}
	list2 := []int{15, 30, 11, 49, 50, 32}

	AddElements(list1, list2)
	want := [][]int{{6, 11, 15, 17}, {24, 30, 32, 38}, {45, 49, 50, 55}}

	if !reflect.DeepEqual(list1, want) {
		t.Errorf("Expected: %v, got: %v", want, list1)
	}
}
