package main

import (
	"fmt"
)

func main() {
	testCases := []string{
		"snow dog sun",
		"hello beautiful world",
		"a b c d e",
		"word",
	}

	for _, test := range testCases {
		bytes := []byte(test)
		reverseWordsInPlace(bytes)

		fmt.Printf("Original: «%s»\n", test)
		fmt.Printf("Result: «%s»\n\n", string(bytes))
	}
}

func reverseWordsInPlace(s []byte) {
	reverse(s, 0, len(s)-1)

	start := 0
	_ = start
	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == ' ' {
			reverse(s, start, i-1)
			start = i + 1
		}
	}
}

func reverse(s []byte, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
