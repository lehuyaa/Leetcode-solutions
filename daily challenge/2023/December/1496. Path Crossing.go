package main

import (
	"fmt"
	"strconv"
)

func isPathCrossing(path string) bool {
	x, y := 0, 0
	visited := make(map[string]bool)
	visited["0-0"] = true
	for i := 0; i < len(path); i++ {
		if path[i] == 'N' {
			y++
		} else if path[i] == 'E' {
			x++
		} else if path[i] == 'S' {
			y--
		} else {
			x--
		}
		_, ok := visited[strconv.Itoa(x)+"-"+strconv.Itoa(y)]
		if ok {
			return true
		}
		visited[strconv.Itoa(x)+"-"+strconv.Itoa(y)] = true
	}

	return false
}

func main() {
	s1 := "NNSWWEWSSESSWENNW"
	s2 := "NESWW"
	fmt.Println("TEST CASE 1", isPathCrossing(s1) == true)
	fmt.Println("TEST CASE 2", isPathCrossing(s2) == true)

}
