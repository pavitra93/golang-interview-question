package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}

func longestCommonPrefix(strs []string) string {
	commonPrefix := ""
	sort.Strings(strs)
	fmt.Println(strs)
	firstString := strs[0]
	lastString := strs[len(strs)-1]
	//for i := 0; i < len(firstString); i++ {
	//	for j := 1; j < len(strs); j++ {
	//		if firstString[i] != strs[j][i] {
	//			return commonPrefix
	//		}
	//	}
	//	commonPrefix += string(firstString[i])
	//
	//}

	for i := 0; i < len(firstString); i++ {
		if firstString[i] != lastString[i] {
			return commonPrefix
		}
		commonPrefix += string(firstString[i])
	}

	return commonPrefix
}
