package main

func mySqrt(x int) int {
	start := 1
	end := x
	for start <= end {
		mid := (start + end) / 2

		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return end
}
