package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	size := len(cost)
	if size == 2 {
		return min(cost[0], cost[1])
	}

	dp := make([]int, size+1)
	dp[0] = 0
	dp[1] = 0

	for i := 2; i <= size; i++ {
		dp[i] = min(dp[i-2]+cost[i-2], dp[i-1]+cost[i-1])
	}

	return dp[size]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
}
