package main

import "fmt"

func maxScore(s string) int {
	scoreZero := []int{}
	maxZero := 0
	if s[0] == '0' {
		scoreZero = append(scoreZero, 1)
		maxZero++
	} else {
		scoreZero = append(scoreZero, 0)
	}

	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			scoreZero = append(scoreZero, scoreZero[i-1]+1)
			maxZero++
		} else {
			scoreZero = append(scoreZero, scoreZero[i-1])
		}
	}

	maxScore := 0
	for i := 0; i < len(s)-1; i++ {
		left := scoreZero[i]
		right := len(s) - i - 1 - maxZero + scoreZero[i]
		if left+right > maxScore {
			maxScore = left + right
		}
	}
	return maxScore
}

func main() {
	s1 := "011101"
	s2 := "00111"
	s3 := "1111"
	fmt.Println("TEST CASE 1", maxScore(s1) == 5)
	fmt.Println("TEST CASE 2", maxScore(s2) == 5)
	fmt.Println("TEST CASE 3", maxScore(s3) == 3)

}
