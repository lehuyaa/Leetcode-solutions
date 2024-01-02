package main

func maxSubArray(nums []int) int {
	size := len(nums)
	if size == 1 {
		return nums[0]
	}
	maxSum := -10001
	sum := 0

	for i := 0; i < size; i++ {
		sum += nums[i]
		if sum > maxSum {
			maxSum = sum
		}
		if sum < 0 {
			sum = 0
		}

	}

	return maxSum
}
