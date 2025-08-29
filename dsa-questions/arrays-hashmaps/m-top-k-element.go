package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}

func topKFrequent(nums []int, k int) []int {
	topKFrequent := make([]int, k)
	FrequencyMap := make(map[int]int)
	type KV struct {
		key   int
		value int
	}

	for _, num := range nums {
		FrequencyMap[num]++
	}

	var freqSlice []KV
	for k, v := range FrequencyMap {
		freqSlice = append(freqSlice, KV{k, v})
	}

	sort.Slice(freqSlice, func(i, j int) bool {
		return freqSlice[i].value > freqSlice[j].value
	})

	for i := 0; i < k; i++ {
		topKFrequent[i] = freqSlice[i].key
	}

	return topKFrequent
}
