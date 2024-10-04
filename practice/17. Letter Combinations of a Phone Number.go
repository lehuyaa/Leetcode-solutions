package main

import "fmt"

var phoneMap []string
var ans []string
var current_solution string

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	phoneMap = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	ans = []string{}
	current_solution = ""
	backtrack(0, digits)
	return ans
}

func backtrack(index int, digits string) {
	if len(current_solution) == len(digits) {
		ans = append(ans, current_solution)
		return
	}
	letter := phoneMap[digits[index]-'2']
	for i := 0; i < len(letter); i++ {
		current_solution += string(letter[i])
		backtrack(index+1, digits)
		current_solution = current_solution[:len(current_solution)-1]
	}
}

func main() {
	fmt.Println(letterCombinations("23"))
}
