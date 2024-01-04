package main

func singleNumber(nums []int) int {
	mask := 0
	for i := 0; i < len(nums); i++ {
		mask ^= nums[i]
	}
	return mask
}
