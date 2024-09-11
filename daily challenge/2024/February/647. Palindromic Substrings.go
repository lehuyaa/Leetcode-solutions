package main

func countSubstrings(s string) int {
	size := len(s)
	if size == 0 {
		return 0
	}
	if size == 1 {
		return 1
	}
	check := make([][]int, size)
	for i := 0; i < size; i++ {
		check[i] = make([]int, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size-i; j++ {
			if i == 0 {
				check[j][j+i] = 1
			} else if i == 1 {
				if s[j] == s[j+i] {
					check[j][j+i] = 1
				}
			} else {
				if s[j] == s[j+i] && check[j+1][j+i-1] == 1 {
					check[j][j+i] = 1
				}
			}
		}
	}
	result := 0
	for i := 0; i < size; i++ {
		for j := i; j < size; j++ {
			if check[i][j] == 1 {
				result++
			}
		}
	}

	return result
}
