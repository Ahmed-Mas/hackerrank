package main

import (
	"fmt"
)

func flippingBits(n int64) int64 {
	// Write your code here
	return int64(^uint32(n))
}

func main() {
	fmt.Println(flippingBits(2147483647))
}
