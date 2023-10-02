package main

import (
	"reflect"
	"testing"
)

func TestFindAll(t *testing.T) {
	data := map[int][]float64{
		1: []float64{3.8, 4.6, 5.2},
		2: []float64{2.2, 3.5, 4.9},
		3: []float64{2.7, 3.1, 4.1},
	}

	testCases := []struct {
		name      string
		data      map[int][]float64
		keys      []int
		operator  string
		targetNum float64
		want      map[int][]float64
	}{
		{
			name:      "Case 1",
			data:      data,
			keys:      []int{1, 3},
			operator:  "G",
			targetNum: 4.0,
			want:      map[int][]float64{1: {4.6, 5.2}, 3: {4.1}},
		},
		{
			name:      "Case 2",
			data:      data,
			keys:      []int{2, 3},
			operator:  "G",
			targetNum: 3.0,
			want:      map[int][]float64{2: {3.5, 4.9}, 3: {3.1, 4.1}},
		}, {
			name:      "Case 3",
			data:      data,
			keys:      []int{1, 2, 3},
			operator:  "LE",
			targetNum: 3.5,
			//I think this example has mistake, because text of exercise says not to put key with no found values
			//want:      map[int][]float64{1: {}, 2: {2.2, 3.5}, 3: {2.7, 3.1}},
			want: map[int][]float64{2: {2.2, 3.5}, 3: {2.7, 3.1}},
		},
	}

	t.Parallel()
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got := FindAll(test.data, test.keys, test.operator, test.targetNum)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Expected: %v, got: %v", test.want, got)
			}
		})
	}
}
