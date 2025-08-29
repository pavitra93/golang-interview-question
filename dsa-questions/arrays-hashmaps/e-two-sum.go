package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	complimentMap := make(map[int]int)
	for i, num := range nums {
		if val, ok := complimentMap[target-num]; ok {
			return []int{val, i}
		}

		complimentMap[num] = i

	}

	return nil

}
