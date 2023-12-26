package main

import "fmt"

func numRollsToTarget(n int, k int, target int) int {
	MOD := 1000000007
	if target == 0 {
		return 1
	}
	if target < n || n*k < target {
		return 0
	}

	dp := [][]int{}
	for i := 0; i < n+1; i++ {
		row := []int{}
		for j := 0; j <= target; j++ {
			row = append(row, 0)
		}
		dp = append(dp, row)
	}
	dp[0][0] = 1
	for i := 1; i <= target; i++ {
		if i > k {
			dp[1][i] = 0
		} else {
			dp[1][i] = 1
		}
	}

	for i := 2; i <= n; i++ {
		for j := 1; j <= target; j++ {
			for x := 1; x <= k; x++ {
				if j-x > 0 {
					dp[i][j] = (dp[i][j] + dp[i-1][j-x]) % MOD

				}
			}
		}
	}

	return dp[n][target]
}

func main() {
	n1, k1, t1 := 1, 6, 3
	n2, k2, t2 := 2, 6, 7
	n3, k3, t3 := 30, 30, 500
	fmt.Println("TEST CASE 1", numRollsToTarget(n1, k1, t1) == 1)
	fmt.Println("TEST CASE 2", numRollsToTarget(n2, k2, t2) == 6)
	fmt.Println("TEST CASE 3", numRollsToTarget(n3, k3, t3), numRollsToTarget(n3, k3, t3) == 222616187)

}
