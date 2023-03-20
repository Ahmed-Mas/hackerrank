package main

import "fmt"

func rotateLeft(d int32, arr []int32) []int32 {
	// Write your code here
	new := make([]int32, 0, len(arr))
	for i := d; int(i) < len(arr); i += 1 {
		new = append(new, arr[i])
	}

	for i := 0; int32(i) < d; i += 1 {
		new = append(new, arr[i])
	}
	return new
}

func main() {
	fmt.Println(rotateLeft(2, []int32{1, 2, 3, 4, 5}))
}
