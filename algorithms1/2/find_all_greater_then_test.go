package main

import (
	"reflect"
	"testing"
)

func TestFindAllGreaterThen(t *testing.T) {
	numbers := [][]float64{{3.2, 5.4}, {6.3, 1.4}, {2.5, 6.5}}
	testCases := []struct {
		name      string
		data      [][]float64
		targetNum float64
		want      []Message
	}{
		{
			// I couldn't think of better name, because we are just changing 2nd parameter
			name:      "Dataset 1",
			data:      numbers,
			targetNum: 4.8,
			want:      []Message{{5.4, 1, 2}, {6.3, 2, 1}, {6.5, 3, 2}},
		},
		{
			name:      "Dataset 2",
			data:      numbers,
			targetNum: 5.5,
			want:      []Message{{6.3, 2, 1}, {6.5, 3, 2}},
		},
	}

	t.Parallel()
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got := FindAllGreaterThen(test.data, test.targetNum)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("Expected: %v, got: %v", test.want, got)
			}
		})
	}
}
