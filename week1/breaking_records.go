package main

import "fmt"

func breakingRecords(scores []int32) []int32 {
	// Write your code here
	max, min := scores[0], scores[0]
	max_ch, min_ch := int32(0), int32(0)
	for _, v := range scores[1:] {
		if v > max {
			max = v
			max_ch += 1
		}
		if v < min {
			min = v
			min_ch += 1
		}
	}
	return []int32{max_ch, min_ch}
}

func main() {
	fmt.Println(breakingRecords([]int32{12, 24, 10, 24}))
}
