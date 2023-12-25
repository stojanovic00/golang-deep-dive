package main

func AddElements(list1 [][]int, list2 []int) {
	for _, val := range list2 {
		for idx, list := range list1 {
			if val > list[0] && val < list[len(list)-1] {
				InsertSorted(&list1[idx], val)
				break
			}
		}
	}
}

// InsertSorted
// When adding elements to slice, pointer to slices header must be passed
func InsertSorted(array *[]int, num int) {
	firstGreaterNumIndex := len(*array)
	for i := 0; i < len(*array); i++ {
		if num == (*array)[i] {
			//Not inserting duplicates
			return
		} else if num < (*array)[i] {
			firstGreaterNumIndex = i
			break
		}
	}
	*array = append(*array, 0)
	copy((*array)[firstGreaterNumIndex+1:], (*array)[firstGreaterNumIndex:])
	(*array)[firstGreaterNumIndex] = num
}
