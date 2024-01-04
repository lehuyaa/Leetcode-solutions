package main

import (
	"math"
)

func thirdMax(nums []int) int {
	first := math.MinInt
	second := math.MinInt
	third := math.MinInt

	for i := 0; i < len(nums); i++ {
		if nums[i] == first || nums[i] == second || nums[i] == third {
			continue
		}
		if nums[i] > first {
			third = second
			second = first
			first = nums[i]
		} else if nums[i] > second {
			third = second
			second = nums[i]
		} else if nums[i] > third {
			third = nums[i]
		}
	}

	if third == math.MinInt {
		return first
	}

	return third
}
