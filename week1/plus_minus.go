package main

import "fmt"

func plusMinus(arr []int32) {
	// Write your code here
	arr_len := float64(len(arr))
	pos, neg, zero := 0., 0., 0.
	for _, v := range arr {
		if v > 0 {
			pos += 1
		} else if v == 0 {
			zero += 1
		} else {
			neg += 1
		}
	}
	fmt.Printf("%.6f\n%.6f\n%.6f\n", pos/arr_len, neg/arr_len, zero/arr_len)
}

func main() {
	plusMinus([]int32{1, 2, -1, 0, 0, 5, -5})
}
