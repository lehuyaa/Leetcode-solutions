package main

import "fmt"

func checkSubarraySum(nums []int, k int) bool {
	if len(nums) <= 1 {
		return false
	}
	total := 0
	count := make(map[int]int)
	count[0] = -1
	for i := 0; i < len(nums); i++ {
		total += nums[i]
		v, ok := count[total%k]
		if ok {
			if i-v > 1 {
				return true
			}
		} else {
			count[total%k] = i
		}
	}

	return false
}

func main() {
	fmt.Println(checkSubarraySum([]int{23, 2, 4, 6, 6}, 7))
}
