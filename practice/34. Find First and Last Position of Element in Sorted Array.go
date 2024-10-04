package main

func searchRange(nums []int, target int) []int {

	return []int{searchStartPos(nums, target), searchEndPos(nums, target)}
}

func searchStartPos(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	ans := -1
	if len(nums) == 0 {
		return ans
	}
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			ans = mid
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}

func searchEndPos(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	ans := -1
	if len(nums) == 0 {
		return ans
	}
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			ans = mid
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}
