package main

func AboveDiagonalSum(matrix [][]int) int {
	sum := 0
	for idx, row := range matrix {
		for _, num := range row[idx+1:] {
			sum += num
		}
	}

	return sum
}
func BelowDiagonalSum(matrix [][]int) int {
	sum := 0
	for idx, row := range matrix {
		for _, num := range row[:idx] {
			sum += num
		}
	}

	return sum
}
