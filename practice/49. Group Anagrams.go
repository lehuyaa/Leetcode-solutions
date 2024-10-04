package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{strs}
	}
	ans := [][]string{}
	group := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		key := getKey(strs[i])
		v, ok := group[key]
		if ok {
			v = append(v, strs[i])
			group[key] = v
		} else {
			group[key] = []string{strs[i]}
		}
	}

	for _, groupItem := range group {
		ans = append(ans, groupItem)
	}
	return ans
}

func getKey(s string) string {
	count := make([]int, 26)
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}
	key := []byte{}
	for i := 0; i < len(count); i++ {
		countItem := count[i]
		if countItem > 0 {
			for j := 1; j <= countItem; j++ {
				key = append(key, byte(i)+'a')
			}
		}

	}

	return string(key)
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
