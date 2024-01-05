package main

func lengthOfLIS(nums []int) int {
	stack := []int{}
	for i := 0; i < len(nums); i++ {
		if i == 0 || stack[len(stack)-1] < nums[i] {
			stack = append(stack, nums[i])
		} else {
			index := findIndex(stack, nums[i])
			stack[index] = nums[i]
		}
	}

	return len(stack)
}

func findIndex(stack []int, num int) int {
	left, right := 0, len(stack)-1
	for left <= right {
		mid := left + (right-left)/2
		if stack[mid] == num {
			return mid
		} else if stack[mid] > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
