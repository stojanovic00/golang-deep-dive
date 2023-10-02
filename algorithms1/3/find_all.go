package main

func SatisfiesRequest(num, targetNum float64, operator string) bool {
	switch operator {
	case "L":
		return num < targetNum
	case "LE":
		return num <= targetNum
	case "G":
		return num > targetNum
	case "GE":
		return num >= targetNum
	default:
		return false
	}
}

func FindAll(data map[int][]float64, keys []int, operator string, targetNum float64) map[int][]float64 {

	foundNums := map[int][]float64{}

	for _, key := range keys {
		for _, num := range data[key] {
			if SatisfiesRequest(num, targetNum, operator) {
				foundNums[key] = append(foundNums[key], num)
			}
		}
	}

	return foundNums
}
