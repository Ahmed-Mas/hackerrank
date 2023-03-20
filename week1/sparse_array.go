package main

import "fmt"

func matchingStrings(strings []string, queries []string) []int32 {
	// Write your code here
	counts := make(map[string]int32, 0)
	for _, val := range strings {
		counts[val] += 1
	}

	result := make([]int32, 0, len(queries))
	for _, val := range queries {
		result = append(result, counts[val])
	}
	return result
}

func main() {
	fmt.Println(matchingStrings([]string{"ab", "ab", "abc"}, []string{"ab", "abc", "bc"}))
}
