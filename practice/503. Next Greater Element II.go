package main

import "fmt"

func nextGreaterElements(nums []int) []int {
	size := len(nums)
	for i := 0; i < size; i++ {
		nums = append(nums, nums[i])
	}
	stack := []int{}
	result := make([]int, size)
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[i] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		number := nums[i]
		if len(stack) == 0 {
			nums[i] = -1
		} else {
			nums[i] = stack[len(stack)-1]
		}
		stack = append(stack, number)
	}
	for i := 0; i < len(result); i++ {
		result[i] = nums[i]
	}
	return result
}

func main() {
	fmt.Println(nextGreaterElements([]int{1, 2, 1}))
}
