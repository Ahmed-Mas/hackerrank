package main

import (
	"fmt"
	"sort"
)

func missingNumbers(arr []int32, brr []int32) []int32 {
	// Write your code here
	count1 := make(map[int32]int)
	count2 := make(map[int32]int)
	for _, val := range arr {
		count1[val] += 1
	}
	for _, val := range brr {
		count2[val] += 1
	}

	keys := make(map[int32]int)
	for k := range count1 {
		keys[k] = 1
	}
	for k := range count2 {
		keys[k] = 1
	}

	res := make([]int32, 0, 0)
	for k := range keys {
		if count1[k] != count2[k] {
			res = append(res, k)
		}
	}
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	return res
}

func main() {
	fmt.Println(missingNumbers([]int32{7, 2, 5, 3, 5, 3}, []int32{7, 2, 5, 4, 6, 5, 3}))
}
