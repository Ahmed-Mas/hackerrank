package main

import (
	"fmt"
	"sort"
)

func migratoryBirds(arr []int32) int32 {
	// Write your code here
	sort.Slice(arr, func(i int, j int) bool { return arr[i] < arr[j] })
	sightings := make([]int32, arr[len(arr)-1]+1)
	for _, val := range arr {
		sightings[val] += 1
	}
	max_sight, idx := int32(-1), int32(0)
	for i := arr[len(arr)-1]; i >= 0; i -= 1 {
		if sightings[i] >= int32(max_sight) {
			max_sight = sightings[i]
			idx = i
		}
	}
	return idx
}

func main() {
	fmt.Println(migratoryBirds([]int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4}))
	// fmt.Println(migratoryBirds([]int32{1, 4, 4, 4, 5, 3}))
}
