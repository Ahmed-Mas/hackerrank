package main

import (
	"fmt"
	"sort"
)

func sockMerchant(n int32, ar []int32) int32 {
	// Write your code here
	sort.Slice(ar, func(i int, j int) bool { return ar[i] < ar[j] })
	pairs, i := int32(0), 0
	for i < len(ar)-1 {
		if ar[i] == ar[i+1] {
			pairs += 1
			i += 2
		} else {
			i += 1
		}
	}
	return pairs
}

func main() {
	fmt.Println(sockMerchant(7, []int32{1, 2, 1, 2, 1, 3, 2}))
}
