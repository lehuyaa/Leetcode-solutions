package main

import "fmt"

var ans [][]int
var current_solution []int

func combine(n int, k int) [][]int {
	ans = [][]int{}
	current_solution = []int{}
	backtrack(0, k, 1, n)
	return ans
}

func backtrack(index int, k int, start int, n int) {
	if index == k {
		copy_solution := make([]int, len(current_solution))
		copy(copy_solution, current_solution)
		ans = append(ans, copy_solution)
		return
	}
	for i := start; i <= n; i++ {
		current_solution = append(current_solution, i)
		backtrack(index+1, k, i+1, n)
		current_solution = current_solution[:len(current_solution)-1]
	}
}

func main() {
	fmt.Println(combine(4, 2))
}
