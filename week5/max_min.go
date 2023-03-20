package main

import (
	"fmt"
	"sort"
)

func maxMin(k int32, arr []int32) int32 {
	// Write your code here
	sort.Slice(arr, func(i int, j int) bool { return arr[i] < arr[j] })
	fmt.Println(arr)
	min_unfair := arr[k-1] - arr[0]
	for i := 0; i <= len(arr)-int(k); i += 1 {
		fmt.Println(arr[i], arr[i+int(k)-1])
		if arr[i+int(k)-1]-arr[i] < int32(min_unfair) {
			min_unfair = arr[i+int(k)-1] - arr[i]
		}
	}
	return min_unfair
}

func main() {
	fmt.Println(maxMin(3, []int32{100, 200, 300, 350, 400, 401, 402}))
}
