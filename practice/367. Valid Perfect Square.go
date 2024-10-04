package main

func isPerfectSquare(num int) bool {
	left := 1
	right := num

	for left <= right {
		mid := (left + right) / 2

		if mid*mid == num {
			return true
		} else if mid*mid > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}
