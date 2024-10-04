package main

import "fmt"

func subarraySum(nums []int, k int) int {
	if len(nums) == 1 {
		if nums[0] == k {
			return 1
		} else {
			return 0
		}
	}
	ans := 0
	sum := make([]int, len(nums))
	count := make(map[int]int)
	sum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sum[i] = sum[i-1] + nums[i]
	}
	count[0] = 1
	for i := 0; i < len(sum); i++ {
		num := sum[i] - k
		v, ok := count[num]
		if ok {
			ans += v

		}
		count[sum[i]]++
	}

	return ans
}

func main() {
	fmt.Println(subarraySum([]int{1, -1, 0}, 0))
}
