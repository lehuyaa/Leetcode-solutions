package main

import (
	"math"
)

func minOperations(nums []int) int {
	count := 0
	mapCount := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		v, ok := mapCount[nums[i]]
		if ok {
			mapCount[nums[i]] = v + 1
		} else {
			mapCount[nums[i]] = 1
		}
	}
	for _, v := range mapCount {
		if v == 1 {
			return -1
		}
		count += int(math.Ceil(float64(v) / 3.0))
	}

	return count
}
