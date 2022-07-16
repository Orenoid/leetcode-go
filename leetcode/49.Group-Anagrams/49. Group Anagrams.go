package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	cache := make(map[[26]int][]string, 0)
	for _, str := range strs {
		cnt := [26]int{}
		for i := range str {
			cnt[str[i]-'a']++
		}
		cache[cnt] = append(cache[cnt], str)
	}
	var result [][]string
	for _, group := range cache {
		result = append(result, group)
	}
	return result
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{"a"}))
}
