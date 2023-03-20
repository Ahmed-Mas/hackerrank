package main

import (
	"fmt"
)

func caesarCipher(s string, k int32) string {
	// Write your code here
	res := ""
	for _, char := range s {
		if 65 <= char && char <= 90 {
			if char+k > 90 {
				res += string((char+k)%91 + 65)
			} else {
				res += string(char + k)
			}
		} else if 97 <= char && char <= 122 {
			if char+k > 122 {
				res += string((char+k)%123 + 97)
			} else {
				res += string(char + k)
			}
		} else {
			res += string(char)
		}
	}
	return res
}

func main() {
	fmt.Println(caesarCipher("www.abc.xy", 87))
}
