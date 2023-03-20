package main

import "fmt"

func pageCount(n int32, p int32) int32 {
	// Write your code here
	var i, j int32
	if n%2 == 0 {
		i, j = 0, n+1
	} else {
		i, j = 0, n
	}

	flips := int32(0)
	for i != j {
		if i == p || i+1 == p || j == p || j-1 == p {
			return flips
		}
		flips += 1
		i += 2
		j -= 2
	}
	return 0
}

func main() {
	fmt.Println(pageCount(5, 3))
}
