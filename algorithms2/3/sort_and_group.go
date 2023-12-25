package main

import (
	"sort"
)

func SortAndGroup(numbers []int) [][]int {

	sort.Ints(numbers)
	sortedAndGrouped := make([][]int, 3)

	for _, num := range numbers {
		switch {
		case num <= 30:
			sortedAndGrouped[0] = append(sortedAndGrouped[0], num)
		case num > 30 && num <= 60:
			sortedAndGrouped[1] = append(sortedAndGrouped[1], num)
		case num > 60 && num <= 100:
			sortedAndGrouped[2] = append(sortedAndGrouped[2], num)
		}
	}

	return sortedAndGrouped
}
