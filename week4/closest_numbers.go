package main

import (
	"fmt"
	"sort"
)

func closestNumbers(arr []int32) []int32 {
	// Write your code here
	sort.Slice(arr, func(i int, j int) bool { return arr[i] < arr[j] })
	res := make([]int32, 0, 0)
	min_diff := arr[1] - arr[0]
	for i := 0; i < len(arr)-1; i += 1 {
		if arr[i+1]-arr[i] < min_diff {
			res = []int32{arr[i], arr[i+1]}
			min_diff = arr[i+1] - arr[i]
		} else if arr[i+1]-arr[i] == min_diff {
			res = append(res, arr[i], arr[i+1])
		}
	}
	return res
}

func main() {
	fmt.Println(closestNumbers([]int32{-20, -3916237, -357920, -3620601, 7374819, -7330761, 30, 6246457, -6461594, 266854, -520, -470}))
}
