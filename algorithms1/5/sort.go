package main

import "sort"

func SortByKey(data map[string]int) map[string]int {

	sorted := make(map[string]int, len(data))
	keys := make([]string, 0, len(data))

	for key, _ := range data {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		sorted[key] = data[key]
	}

	return sorted
}

func SortByValue(data map[string]int) map[string]int {
	reversed := make(map[int]string, len(data))
	reverseKeys := make([]int, 0, len(data))

	// Reversing key and value in map
	for key, value := range data {
		reversed[value] = key
		reverseKeys = append(reverseKeys, value)
	}

	sort.Ints(reverseKeys)

	sorted := make(map[string]int, len(data))
	for _, key := range reverseKeys {
		sorted[reversed[key]] = key
	}

	return sorted
}
