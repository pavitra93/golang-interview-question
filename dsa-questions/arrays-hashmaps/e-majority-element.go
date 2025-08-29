package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
}

func majorityElement(nums []int) int {
	ele := 0
	count := 0

	for _, num := range nums {
		if count == 0 {
			ele = num
		}

		if num == ele {
			count++
		} else {
			count--
		}
	}

	return ele

}
