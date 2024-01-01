package main

import "sort"

func findContentChildren(g []int, s []int) int {
	result := 0

	sort.SliceStable(g, func(i, j int) bool {
		return g[i] < g[j]
	})
	sort.SliceStable(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	indexG, indexS := 0, 0
	lenG, lenS := len(g), len(s)

	for indexG < lenG && indexS < lenS {
		if g[indexG] <= s[indexS] {
			result++
			indexG++
			indexS++
		} else {
			indexS++
		}
	}

	return result
}
