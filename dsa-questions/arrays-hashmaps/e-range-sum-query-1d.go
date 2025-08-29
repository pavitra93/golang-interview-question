package main

func main() {
	Constructor1D([]int{-2, 0, 3, -5, 2, -1})
	//param_1 := obj.SumRange(left, right)
}

type NumArray struct {
	inputArray []int
	prefixSums []int
}

func Constructor1D(nums []int) NumArray {
	prefixSums := make([]int, len(nums))
	prefixSums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefixSums[i] = prefixSums[i-1] + nums[i]
	}

	return NumArray{nums, prefixSums}
}

func (numArray *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return numArray.prefixSums[right]
	} else {
		return numArray.prefixSums[right] - numArray.prefixSums[left-1]
	}
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
