package main

import (
	"sort"
)

type Ans struct {
	index int
	size  int
}

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = []int{nums[i]}
	}
	max := Ans{0, 1}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 {
				if len(dp[j])+1 > len(dp[i]) {
					dp[i] = append([]int(nil), dp[j]...)
					dp[i] = append(dp[i], nums[i])
				}
				if len(dp[i]) > max.size {
					max.size = len(dp[i])
					max.index = i
				}
			}
		}
	}
	return dp[max.index]
}
