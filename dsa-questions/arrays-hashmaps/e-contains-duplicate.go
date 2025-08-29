package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
}

func containsDuplicate(nums []int) bool {
	mapDuplicates := make(map[int]int)
	found := false
	for _, num := range nums {
		_, ok := mapDuplicates[num]
		if ok {
			found = true
			break
		}
		mapDuplicates[num]++
	}

	fmt.Println(mapDuplicates)

	return found
}
