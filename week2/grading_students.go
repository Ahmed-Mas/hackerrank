package main

import (
	"fmt"
	"math"
)

func gradingStudents(grades []int32) []int32 {
	// Write your code here
	for i, val := range grades {
		nearest_5 := math.Ceil(float64(val)/5.) * 5
		diff := nearest_5 - float64(val)
		if val%5 == 0 || val < 38 || diff >= 3 {
			continue
		}
		grades[i] = int32(nearest_5)
	}
	return grades
}

func main() {
	fmt.Println(gradingStudents([]int32{84, 29, 57}))
}
