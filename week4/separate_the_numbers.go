package main

import (
	"fmt"
	"math/big"
)

func separateNumbers(s string) {
	// Write your code here
	for i := 1; i < len(s); i += 1 {
		head := s[:i]

		cur_str := head
		prev := new(big.Int)
		prev, _ = prev.SetString(cur_str, 10)

		for len(cur_str) < len(s) {
			nxt := big.NewInt(0).Add(prev, big.NewInt(1))
			cur_str = cur_str + nxt.String()
			prev = nxt
		}
		if cur_str == s {
			fmt.Printf("YES %v\n", head)
			return
		}
	}
	fmt.Println("NO")
}

func main() {
	separateNumbers("11111111111111111111111111111111")
}
