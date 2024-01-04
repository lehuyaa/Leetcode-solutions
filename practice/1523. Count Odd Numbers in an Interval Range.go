package main

func countOdds(low int, high int) int {
	if low%2 == 0 {
		low++
	}
	if high%2 == 0 {
		high--
	}

	return (high-low)/2 + 1
}
