package main

func findMatrix(nums []int) [][]int {
	result := [][]int{}
	mapCount := make(map[int]int)
	maxRow := 0

	for i := 0; i < len(nums); i++ {
		v, ok := mapCount[nums[i]]
		if ok {
			mapCount[nums[i]] = v + 1
		} else {
			mapCount[nums[i]] = 1
		}
		newV, _ := mapCount[nums[i]]
		if newV > maxRow {
			maxRow = newV
		}
	}
	for i := 1; i <= maxRow; i++ {
		row := []int{}
		for k, v := range mapCount {
			if v > 0 {
				row = append(row, k)
				mapCount[k]--
			}

		}

		result = append(result, row)
	}
	return result
}
