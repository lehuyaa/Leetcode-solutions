package main

import "sort"

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	jobs := make([][]int, n)
	for i := 0; i < n; i++ {
		jobs[i] = []int{startTime[i], endTime[i], profit[i]}
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][1] < jobs[j][1]
	})

	dp := [][]int{{0, 0}}

	for _, job := range jobs {
		start, end, prof := job[0], job[1], job[2]
		i := bisec(dp, start)
		newProfit := dp[i-1][1] + prof
		if newProfit > dp[len(dp)-1][1] {
			dp = append(dp, []int{end, newProfit})
		}
	}
	return dp[len(dp)-1][1]
}

func bisec(arr [][]int, target int) int {
	low, high := 0, len(arr)

	for low < high {
		mid := (low + high) / 2
		if arr[mid][0] > target {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}
