package main

import "testing"

func TestAboveDiagonalSum(t *testing.T) {
	matrix := [][]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}}

	got := AboveDiagonalSum(matrix)
	want := 40

	if got != want {
		t.Errorf("Expected sum: %v, got: %v", want, got)
	}
}

func TestBelowDiagonalSum(t *testing.T) {
	matrix := [][]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}}

	got := BelowDiagonalSum(matrix)
	want := 20

	if got != want {
		t.Errorf("Expected sum: %v, got: %v", want, got)
	}
}
