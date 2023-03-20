package main

import "fmt"

func countingValleys(steps int32, path string) int32 {
	// Write your code here
	level, valleys := 0, 0
	for _, val := range path {
		prev := level
		if val == 'D' {
			level -= 1
		} else {
			level += 1
		}
		if prev == 0 && level < 0 {
			valleys += 1
		}
	}
	return int32(valleys)
}

func main() {
	fmt.Println(countingValleys(12, "UDDDUDUUDDUU"))
}
