package main

import "fmt"

func search(nums []int, target int) int {
	indexMin := findIndexMin(nums)
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	if len(nums) == 2 {
		if nums[0] == target {
			return 0
		} else if nums[1] == target {
			return 1
		}

		return -1
	}

	var left, right int
	if target > nums[len(nums)-1] {
		left = 0
		right = indexMin - 1
	} else {
		left = indexMin
		right = len(nums) - 1
	}
	if left > right {
		return -1
	}
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func findIndexMin(nums []int) int {
	size := len(nums)
	if size == 1 {
		return nums[0]
	}

	if size == 2 {
		if nums[0] == getMin(nums[0], nums[1]) {
			return 0
		}
		return 1
	}
	left := 0
	right := size - 1
	for left <= right {
		mid := (left + right) / 2
		if (mid == left || nums[mid] < nums[mid-1]) && (mid == right || nums[mid] < nums[mid+1]) {
			return mid
		} else if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return nums[left]
}

func getMin(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	fmt.Println(search([]int{3, 5, 1}, 5))
}
