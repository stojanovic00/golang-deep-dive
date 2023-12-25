package main

func FindMax(array []int) int {
	maxNum := array[0]

	for i := 1; i < len(array); i++ {
		if array[i] > maxNum {
			maxNum = array[i]
		}
	}

	return maxNum
}

func MaxInArray(arrays [][]int) []int {
	maxes := make([]int, 0, len(arrays))

	for _, array := range arrays {
		maxes = append(maxes, FindMax(array))
	}

	return maxes
}
