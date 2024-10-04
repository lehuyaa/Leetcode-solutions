package main

import (
	"fmt"
	"sort"
)

func frequencySort(s string) string {
	freq := make([]int, 255)
	ans := []byte(s)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}
	sort.Slice(ans, func(i int, j int) bool {
		if freq[ans[i]] == freq[ans[j]] {
			return ans[i] > ans[j]
		}
		return freq[ans[i]] > freq[ans[j]]
	})
	return string(ans)
}

func main() {
	fmt.Println(frequencySort("tree"))
}
