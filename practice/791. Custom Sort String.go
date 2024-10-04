package main

import "sort"

func customSortString(order string, s string) string {
	orderMap := make(map[byte]int)
	for i := 0; i < len(order); i++ {
		orderMap[order[i]] = len(order) - i - 1
	}
	arr := []byte(s)
	sort.Slice(arr, func(i, j int) bool {
		indexI, mapI := orderMap[arr[i]]
		indexJ, mapJ := orderMap[arr[j]]
		if mapI && !mapJ {
			return true
		} else if !mapI && mapJ {
			return false
		} else if !mapI && !mapJ {
			return true
		}
		return indexI > indexJ
	})

	return string(arr)
}
