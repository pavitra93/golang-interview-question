package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

// helper: returns sorted version of a string
func sortString(s string) string {
	chars := []rune(s)
	sort.Slice(chars, func(i, j int) bool { return chars[i] < chars[j] })
	return string(chars)
}

func groupAnagrams(strs []string) [][]string {
	grouped := make(map[string][]string)

	for _, s := range strs {
		key := sortString(s) // sorted string as key
		grouped[key] = append(grouped[key], s)
	}

	// collect only values
	res := make([][]string, 0, len(grouped))
	for _, v := range grouped {
		res = append(res, v)
	}

	return res
}
