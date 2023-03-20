package main

import "fmt"

func lonelyinteger(a []int32) int32 {
	// Write your code here
	n := int32(0)
	for _, val := range a {
		n = n ^ val
	}
	return n
}

func main() {
	fmt.Println(lonelyinteger([]int32{1, 2, 3, 4, 3, 2, 1}))
}
