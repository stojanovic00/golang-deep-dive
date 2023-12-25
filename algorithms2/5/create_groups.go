package main

import (
	"fmt"
	"sort"
)

func CreateGroups(list []int, lowBound, highBound, separator1, separator2 int) ([][]int, error) {

	if lowBound >= highBound {
		return nil, fmt.Errorf("low bound greater or equal then high")
	}

	if separator1 >= separator2 {
		return nil, fmt.Errorf("low separator greater or equal then high")
	}

	sort.Ints(list)

	if separator1 < list[0] {
		return nil, fmt.Errorf("low separator out of bounds")
	}
	if separator2 > list[len(list)-1] {
		return nil, fmt.Errorf("high separator out of bounds")
	}

}
