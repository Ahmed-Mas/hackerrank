package main

import "fmt"

func marsExploration(s string) int32 {
	// Write your code here
	msg := "SOS"
	changed := 0
	for i := 0; i < len(s); i += 1 {
		if s[i] != msg[i%3] {
			changed += 1
		}
	}
	return int32(changed)
}

func main() {
	fmt.Println(marsExploration("SOSTOT"))
}
