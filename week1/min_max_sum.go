package main

import (
	"fmt"
	"sort"
)

func miniMaxSum(arr []int32) {
	// Write your code here
	sort.Slice(arr, func(i int, j int) bool { return arr[i] < arr[j] })
	sum := int64(0)
	for _, v := range arr[1:4] {
		sum += int64(v)
	}

	fmt.Printf("%d %d", sum+int64(arr[0]), sum+int64(arr[4]))
}

func main() {
	miniMaxSum([]int32{256741038, 623958417, 467905213, 714532089, 938071625})
}
