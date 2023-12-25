package main

import (
	"reflect"
	"testing"
)

func TestInsertSorted(t *testing.T) {
	data := []int{1, 2, 5, 7, 8, 11, 14}

	testCases := []struct {
		Name string
		Num  int
		Want []int
	}{
		{
			Name: "Case 1",
			Num:  4,
			Want: []int{1, 2, 4, 5, 7, 8, 11, 14},
		},
		{
			Name: "Case 2",
			Num:  9,
			Want: []int{1, 2, 5, 7, 8, 9, 11, 14},
		},
		{
			Name: "Case 3",
			Num:  12,
			Want: []int{1, 2, 5, 7, 8, 11, 12, 14},
		},
	}

	t.Parallel()

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {

			got := InsertSorted(data, test.Num)

			if !reflect.DeepEqual(got, test.Want) {
				t.Errorf("Exptected %v, got: %v", test.Want, got)
			}
		})
	}

}
