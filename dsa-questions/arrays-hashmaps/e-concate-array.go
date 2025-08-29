package main

import "fmt"

func main() {

	fmt.Printf("%v", getConcatenation([]int{1, 2, 3}))
}

func getConcatenation(nums []int) []int {
	concatNums := make([]int, len(nums)*2)
	concatNums = append(concatNums, nums...)
	concatNums = append(concatNums, nums...)
	return concatNums

}
