package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	index := 1
	start := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] != start {
			nums[index] = nums[i]
			start = nums[i]
			index++
		}
	}

	return index
}

func main() {
	nums1 := []int{1, 1, 2}
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	fmt.Println("TEST CASE 1", removeDuplicates(nums1), removeDuplicates(nums1) == 2)
	fmt.Println("TEST CASE 2", removeDuplicates(nums2), removeDuplicates(nums2) == 5)

}
