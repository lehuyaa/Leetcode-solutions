package main

func sortColors(nums []int) {
	mapCount := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			mapCount[0]++
		}
		if nums[i] == 1 {
			mapCount[1]++
		}
		if nums[i] == 2 {
			mapCount[2]++
		}
	}
	for i := 0; i < len(nums); i++ {
		if i < mapCount[0] {
			nums[i] = 0
		} else if i >= mapCount[0] && i < mapCount[0]+mapCount[1] {
			nums[i] = 1
		} else {
			nums[i] = 2
		}
	}
}

func main() {
	sortColors([]int{2, 0, 2, 1, 1, 0})
}
