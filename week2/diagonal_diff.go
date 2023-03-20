package main

import (
	"fmt"
	"math"
)

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	right_diag, left_diag := int32(0), int32(0)
	for idx := 0; idx < len(arr); idx += 1 {
		right_diag += arr[idx][idx]
		left_diag += arr[idx][len(arr)-1-idx]
	}
	return int32(math.Abs(float64(right_diag) - float64(left_diag)))
}

func main() {
	fmt.Println(diagonalDifference([][]int32{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}))
}
