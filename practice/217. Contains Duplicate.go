package main

func containsDuplicate(nums []int) bool {
	visited := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		_, ok := visited[nums[i]]
		if ok {
			return true
		} else {
			visited[nums[i]] = 1
		}
	}

	return false
}
