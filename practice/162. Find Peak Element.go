package main

import "fmt"

func findPeakElement(nums []int) int {
	size := len(nums)
	if size == 1 {
		return 0
	}
	if size == 2 {
		if nums[0] > nums[1] {
			return 0
		}
		return 1
	}

	left := 0
	right := size - 1

	for left <= right {
		mid := (left + right) / 2
		if (mid == 0 || nums[mid] > nums[mid-1]) && (mid == size-1 || nums[mid] > nums[mid+1]) {
			return mid
		} else if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
}
