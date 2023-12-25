package main

import (
	"fmt"
)

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}

	n := len(s)
	dp := make([]int, n+1)
	dp[n] = 1
	if s[n-1] != '0' {
		dp[n-1] = 1
	}

	for i := n - 2; i >= 0; i-- {
		if s[i] != '0' {
			dp[i] += dp[i+1]
		}
		if s[i] == '1' || s[i] == '2' && s[i+1] < '7' {
			dp[i] += dp[i+2]
		}
	}
	return dp[0]
}

func main() {
	s1 := "12"
	s2 := "226"
	s3 := "06"
	fmt.Println("TEST CASE 1", numDecodings(s1) == 2)
	fmt.Println("TEST CASE 2", numDecodings(s2) == 3)
	fmt.Println("TEST CASE 3", numDecodings(s3) == 0)

}
