package main

import (
	"fmt"
	"strings"
)

func pangrams(s string) string {
	// Write your code here
	s = strings.ToLower(s)
	chars := make([]int32, 26)
	for _, val := range s {
		if val == ' ' {
			continue
		}
		chars[rune(val)%97] += 1
	}
	for _, val := range chars {
		if val == 0 {
			return "not pangram"
		}
	}
	return "pangram"
}

func main() {
	fmt.Println(pangrams("We promptly judged antique ivory buckles for the next prize"))
}
