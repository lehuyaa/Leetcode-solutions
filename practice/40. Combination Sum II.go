package main

import (
	"fmt"
	"sort"
)

var ans [][]int
var current_solution []int
var current_sum int

func combinationSum2(candidates []int, target int) [][]int {
	ans = [][]int{}
	current_solution = []int{}
	current_sum = 0
	sort.Ints(candidates)
	backtrack(0, candidates, target)
	return ans
}

func backtrack(index int, candidates []int, target int) {
	if current_sum > target {
		return
	}
	if current_sum == target {
		copy_solution := make([]int, len(current_solution))
		copy(copy_solution, current_solution)
		ans = append(ans, copy_solution)
		return
	}

	for i := index; i < len(candidates); i++ {
		if i > index && candidates[i] == candidates[i-1] {
			fmt.Println(candidates[i])
			continue
		}
		current_solution = append(current_solution, candidates[i])
		current_sum += candidates[i]
		backtrack(i+1, candidates, target)
		current_sum -= candidates[i]
		current_solution = current_solution[:len(current_solution)-1]
	}
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
