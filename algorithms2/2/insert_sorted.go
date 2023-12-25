package main

func InsertSorted(array []int, num int) []int {
	firstGreaterNumIndex := len(array)
	for i := 0; i < len(array); i++ {
		if num < array[i] {
			firstGreaterNumIndex = i
			break
		}
	}
	array = append(array, 0)
	copy(array[firstGreaterNumIndex+1:], array[firstGreaterNumIndex:])
	array[firstGreaterNumIndex] = num
	return array
}
