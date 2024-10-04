package main

func findMin(nums []int) int {
	size := len(nums)
	if size == 1 {
		return nums[0]
	}

	if size == 2 {
		return getMin(nums[0], nums[1])
	}

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if (mid == left || nums[mid-1] > nums[mid]) && (mid == right || nums[mid+1] > nums[mid]) {
			return nums[mid]
		} else if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return nums[left]
}

func getMin(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
