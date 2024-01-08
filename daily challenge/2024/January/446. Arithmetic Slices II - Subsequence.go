package main

func numberOfArithmeticSlices(nums []int) int {
	count := 0
	dp := make([]map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		dp[i] = make(map[int]int)
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]
			if val, found := dp[j][diff]; found {
				dp[i][diff] += val + 1
				count += val
			} else {
				dp[i][diff] += 1
			}
		}
	}
	return count
}
