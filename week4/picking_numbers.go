package main

import (
	"fmt"
	"math"
	"sort"
)

func pickingNumbers(a []int32) int32 {
	// Write your code here
	sort.Slice(a, func(i int, j int) bool { return a[i] < a[j] })
	cur_longest, s, e := int32(0), 0, 1
	for e < len(a) {
		temp_l := 0
		for e < len(a) && math.Abs(float64(a[e]-a[s])) <= 1 {
			temp_l += 1
			e += 1
		}
		if temp_l > int(cur_longest) {
			cur_longest = int32(temp_l)
		}
		s = e
		e += 1
	}
	return cur_longest + 1
}

func main() {
	fmt.Println(pickingNumbers([]int32{1, 2, 2, 3, 1, 2}))
}
