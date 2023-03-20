package main

import (
	"fmt"
	"sort"
)

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	// Write your code here
	is_divis := int32(0)
	sort.Slice(ar, func(i int, j int) bool { return ar[i] > ar[j] })
	for idx, val := range ar {
		for _, val2 := range ar[idx+1:] {
			if (val+val2)%k == 0 {
				is_divis += 1
			}
		}
	}
	return is_divis
}

func main() {
	fmt.Println(divisibleSumPairs(6, 3, []int32{1, 3, 2, 6, 1, 2}))
}
