package main

import (
	"fmt"
	"sort"
)

func maxWidthOfVerticalArea(points [][]int) int {
	maxWidest := 0

	sort.SliceStable(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	for i := 1; i < len(points); i++ {
		if points[i][0]-points[i-1][0] > maxWidest {
			maxWidest = points[i][0] - points[i-1][0]
		}
	}
	return maxWidest
}

func main() {
	points1 := [][]int{{8, 7}, {9, 9}, {7, 4}, {9, 7}}
	points2 := [][]int{{3, 1}, {9, 0}, {1, 0}, {1, 4}, {5, 3}, {8, 8}}

	fmt.Println("TEST CASE 1", maxWidthOfVerticalArea(points1) == 1)
	fmt.Println("TEST CASE 2", maxWidthOfVerticalArea(points2) == 3)

}
