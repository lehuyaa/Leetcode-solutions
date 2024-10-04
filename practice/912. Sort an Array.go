package main

import "fmt"

func sortArray(nums []int) []int {
	size := len(nums)
	ans := make([]int, size)
	ans = mergeSort(nums)
	return ans
}

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	left := mergeSort(nums[:len(nums)/2])
	right := mergeSort(nums[len(nums)/2:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	ans := []int{}
	l, r := 0, 0
	sizeL := len(left)
	sizeR := len(right)

	for l < sizeL && r < sizeR {
		if left[l] < right[r] {
			ans = append(ans, left[l])
			l++
		} else {
			ans = append(ans, right[r])
			r++
		}
	}
	for l < sizeL {
		ans = append(ans, left[l])
		l++
	}
	for r < sizeR {
		ans = append(ans, right[r])
		r++
	}
	return ans
}

func main() {
	fmt.Println(sortArray([]int{5, 1, 1, 2, 0, 0}))
}
