package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := []int{}
	result := make([]int, len(nums1))
	max := make(map[int]int)

	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums2[i] > nums2[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			max[nums2[i]] = -1
		} else {
			max[nums2[i]] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	for i := 0; i < len(nums1); i++ {
		item := max[nums1[i]]
		if item == -1 {
			result[i] = -1
		} else {
			result[i] = nums2[max[nums1[i]]]
		}
	}

	return result
}
