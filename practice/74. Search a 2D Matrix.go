package main

func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		if binarySearch(matrix[i], target) {
			return true
		}
	}

	return false
}

func binarySearch(arr []int, target int) bool {
	size := len(arr)
	if size == 0 {
		return false
	}

	if arr[0] > target || arr[size-1] < target {
		return false
	}
	left := 0
	right := size - 1

	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return true
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}
