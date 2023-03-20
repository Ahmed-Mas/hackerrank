package main

import (
	"fmt"
	"sort"
)

func twoArrays(k int32, A []int32, B []int32) string {
	// Write your code here
	sort.Slice(A, func(i int, j int) bool { return A[i] < A[j] })
	sort.Slice(B, func(i int, j int) bool { return B[i] > B[j] })
	for i := 0; i < len(A); i += 1 {
		if A[i]+B[i] < k {
			return "NO"
		}
	}
	return "YES"
}

func main() {
	fmt.Println(twoArrays(3, []int32{2, 1, 3}, []int32{7, 8, 9}))
}
