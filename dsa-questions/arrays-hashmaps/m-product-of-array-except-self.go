package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}

func productExceptSelf(nums []int) []int {
	productArray := make([]int, len(nums))
	//for i := 0; i < len(nums); i++ {
	//	product := 1
	//	for j := 0; j < len(nums); j++ {
	//		if i != j {
	//			product *= nums[j]
	//		}
	//	}
	//	productArray[i] = product
	//}

	return productArray
}
