package main

import (
	"fmt"
	"sort"
)

func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)
	size := len(nums)
	sum := make([]int, size)
	sum[0] = nums[0]
	for i := 1; i < size; i++ {
		sum[i] = nums[i] + sum[i-1]
	}
	// sum = 1,3,7,12
	ans := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		ans[i] = search(sum, queries[i])
	}

	return ans
}

func search(nums []int, target int) int {
	size := len(nums)
	left := 0
	right := size - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

func main() {
	fmt.Println(answerQueries([]int{4, 5, 2, 1}, []int{3, 10, 21}))
}
