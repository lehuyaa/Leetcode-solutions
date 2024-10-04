package main

var ans []string
var current_solution string
var open int

func generateParenthesis(n int) []string {
	ans = []string{}
	current_solution = ""
	open = 0
	backtrack(0, n)
	return ans
}

func backtrack(index int, n int) {
	if index == 2*n && open == n {
		ans = append(ans, current_solution)
		return
	}

	if open < n {
		open++
		current_solution += "("
		backtrack(index+1, n)
		current_solution = current_solution[:len(current_solution)-1]
		open--
	}

	if open >= index+1-open {
		current_solution += ")"
		backtrack(index+1, n)
		current_solution = current_solution[:len(current_solution)-1]
	}

}
