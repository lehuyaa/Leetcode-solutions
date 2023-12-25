package main

import (
	"fmt"
	"reflect"
)

func twoSum(nums []int, target int) []int {
	check := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		number := target - nums[i]
		value, prs := check[number]
		if prs {
			return []int{value, i}
		} else {
			check[nums[i]] = i
		}
	}

	return []int{}
}

func main() {
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	nums2 := []int{3, 2, 4}
	target2 := 6
	nums3 := []int{3, 3}
	target3 := 6

	fmt.Println("TEST CASE 1", compareArrays(twoSum(nums1, target1), []int{0, 1}))
	fmt.Println("TEST CASE 2", compareArrays(twoSum(nums2, target2), []int{1, 2}))
	fmt.Println("TEST CASE 3", compareArrays(twoSum(nums3, target3), []int{0, 1}))

}

func compareArrays(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	return reflect.DeepEqual(arr1, arr2)
}
