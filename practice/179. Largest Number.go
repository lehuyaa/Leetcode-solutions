package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	ans := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = strconv.Itoa(nums[i])
	}

	sort.Slice(ans, func(i int, j int) bool {
		return ans[i]+ans[j] > ans[j]+ans[i]
	})

	if ans[0] == "0" {
		return "0"
	}
	return strings.Join(ans, "")
}

func main() {
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
}
