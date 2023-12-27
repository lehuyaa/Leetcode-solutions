package main

type Colors struct {
	value byte
	index int
}

func minCost(colors string, neededTime []int) int {
	stack := []Colors{}
	time := 0
	for i := 0; i < len(colors); i++ {
		if len(stack) == 0 {
			stack = append(stack, Colors{colors[i], i})
			continue
		}
		size := len(stack) - 1
		top := stack[size]
		if top.value != colors[i] {
			stack = append(stack, Colors{colors[i], i})
		} else {
			if neededTime[top.index] > neededTime[i] {
				time += neededTime[i]
			} else {
				time += neededTime[top.index]
				stack = stack[:len(stack)-1]
				stack = append(stack, Colors{colors[i], i})
			}
		}
	}
	return time
}
