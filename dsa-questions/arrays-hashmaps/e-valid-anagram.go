package main

import (
	"fmt"
)

func main() {
	fmt.Println(isAnagram("aacc", "ccac"))
}

func isAnagram(s string, t string) bool {

	isAnagram := true

	if len(s) != len(t) {
		return false
	}

	stringMapLimited := make([]int, 26)
	for _, char := range s {
		stringMapLimited[char-'a']++
	}

	for _, char := range t {
		stringMapLimited[char-'a']--
		if stringMapLimited[char-'a'] < 0 {
			isAnagram = false
			break
		}
	}

	//stringMap := make(map[rune]int)
	//for _, char := range s {
	//	stringMap[char]++
	//}
	//
	//for _, char := range t {
	//	stringMap[char]--
	//	if stringMap[char] < 0 {
	//		isAnagram = false
	//		break
	//	}
	//}
	//
	//fmt.Println(stringMap)
	return isAnagram
}
