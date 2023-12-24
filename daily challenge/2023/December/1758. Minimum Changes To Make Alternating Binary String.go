package main

import (
	"fmt"
)

func minOperations(s string) int {
	start0 := 0
	start1 := 0
	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			if s[i] == '0' {
				start1++
			} else {
				start0++
			}
		} else {
			if s[i] == '0' {
				start0++
			} else {
				start1++
			}
		}
	}
	if start1 > start0 {
		return start0
	}
	return start1
}

func main() {
	s1 := "0100"
	s2 := "10"
	s3 := "1111"
	fmt.Println("TEST CASE 1", minOperations(s1) == 1)
	fmt.Println("TEST CASE 2", minOperations(s2) == 0)
	fmt.Println("TEST CASE 3", minOperations(s3) == 2)

}
