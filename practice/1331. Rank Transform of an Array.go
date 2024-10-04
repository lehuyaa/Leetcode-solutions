package main

import "sort"

func arrayRankTransform(arr []int) []int {
	copyArr := make([]int, len(arr))
	copy(copyArr, arr)
	sort.Ints(arr)

	ans := make([]int, len(arr))
	mapRank := make(map[int]int)
	rank := 1
	for i := 0; i < len(arr); i++ {
		_, ok := mapRank[arr[i]]
		if !ok {
			mapRank[arr[i]] = rank
			rank++
		}
	}

	for i := 0; i < len(copyArr); i++ {
		ans[i] = mapRank[copyArr[i]]
	}

	return ans
}
