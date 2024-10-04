package main

import "fmt"

func pancakeSort(arr []int) []int {
	size := len(arr)
	currentIndex := size - 1
	ans := []int{}
	for currentIndex >= 0 {
		maxIndex := 0
		for i := 0; i <= currentIndex; i++ {
			if arr[i] > arr[maxIndex] {
				maxIndex = i
			}
		}
		arr = reverse(arr, maxIndex)
		ans = append(ans, maxIndex+1)
		arr = reverse(arr, currentIndex)
		ans = append(ans, currentIndex+1)
		currentIndex--
	}

	return ans
}

func reverse(nums []int, index int) []int {
	l, r := 0, index
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}

	return nums
}

func main() {
	fmt.Println(pancakeSort([]int{3, 2, 4, 1}))
}
