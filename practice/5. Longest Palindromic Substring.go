package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	size := len(s)
	dp := [][]int{}
	for i := 0; i < size; i++ {
		row := []int{}
		for j := 0; j < size; j++ {
			if i == j {
				row = append(row, 1)
			} else {
				row = append(row, 0)
			}
		}
		dp = append(dp, row)
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size-i; j++ {
			if i == 0 {
				dp[i][j] = 1
			} else if i == 1 {
				if s[j] == s[j+i] {
					dp[j][j+i] = 1
				} else {
					dp[j][j+i] = 0
				}
			} else {
				if s[j] == s[j+i] && dp[j+1][j+i-1] == 1 {
					dp[j][j+i] = 1
				} else {
					dp[j][j+i] = 0
				}
			}
		}
	}
	start := 0
	end := 0
	maxScore := 0
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if dp[i][j] == 1 && j-i > maxScore {
				maxScore = j - i
				start = i
				end = j
			}
		}
	}
	return s[start : end+1]
}

func main() {
	s1 := "babad"
	s2 := "cbbd"
	fmt.Println("TEST CASE 1", longestPalindrome(s1) == "bab")
	fmt.Println("TEST CASE 2", longestPalindrome(s2) == "bb")

}
