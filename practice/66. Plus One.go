package main

func plusOne(digits []int) []int {
	dis := 1
	size := len(digits)
	for i := size - 1; i >= 0; i-- {
		number := digits[i] + dis
		if number >= 10 {
			digits[i] = number - 10
			dis = 1
		} else {
			digits[i] = number
			dis = 0
		}
	}
	if dis == 1 {
		result := append([]int{1}, digits...)
		return result
	}
	return digits
}
