package main

import (
	"fmt"
	"math"
	"sort"
)

func minimumAbsoluteDifference(arr []int32) int32 {
	// Write your code here
	sort.Slice(arr, func(i int, j int) bool { return arr[i] < arr[j] })
	min_diff := math.Abs(float64(arr[0] - arr[1]))
	for i := 0; i < len(arr)-1; i += 1 {
		temp_diff := math.Abs(float64(arr[i] - arr[i+1]))
		if temp_diff < min_diff {
			min_diff = temp_diff
		}
	}
	return int32(min_diff)
}

func main() {
	fmt.Println(minimumAbsoluteDifference([]int32{-2, 2, 4}))
}
