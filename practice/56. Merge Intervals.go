package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}
	ans := [][]int{}
	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	start := intervals[0][0]
	end := intervals[0][1]
	ans = append(ans, []int{start, end})
	for i := 1; i < len(intervals); i++ {
		s := intervals[i][0]
		e := intervals[i][1]

		if s > end {
			ans = append(ans, []int{s, e})
			start = s
			end = e
		} else if s <= end && e > end {
			end = e
			ans[len(ans)-1][1] = e
		}
	}

	return ans
}

func main() {
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
}
