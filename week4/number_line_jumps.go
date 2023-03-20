package main

import (
	"fmt"
	"math"
)

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// Write your code here
	if x1 == x2 {
		return "YES"
	}

	if math.Abs(float64((x1+v1)-(x2+v2))) >= math.Abs(float64(x1-x2)) {
		return "NO"
	}
	var smaller_x, smaller_v, larger_v, larger_x int32
	if x1 < x2 {
		smaller_x = x1
		larger_x = x2

		smaller_v = v1
		larger_v = v2
	} else {
		smaller_x = x2
		larger_x = x1

		smaller_v = v2
		larger_v = v1
	}

	for smaller_x <= larger_x {
		smaller_x += smaller_v
		larger_x += larger_v
		if smaller_x == larger_x {
			return "YES"
		}
	}
	return "NO"
}

func main() {
	fmt.Println(kangaroo(2, 1, 1, 2))
}
