package main

import "fmt"

func maxOperations(nums []int, k int) int {
	count := make(map[int]int)
	size := len(nums)
	for i := 0; i < size; i++ {
		count[nums[i]]++
	}
	ans := 0
	for i := 0; i < size; i++ {
		cur, _ := count[nums[i]]
		v, ok := count[k-nums[i]]
		if nums[i] != k-nums[i] {
			if cur > 0 && ok && v > 0 {
				count[nums[i]]--
				count[k-nums[i]]--
				ans++
			}
		} else {
			if cur > 1 {
				count[nums[i]]--
				count[k-nums[i]]--
				ans++
			}
		}

	}

	return ans
}

func main() {
	fmt.Println(maxOperations([]int{3, 1, 3, 4, 3}, 6))
}
