package main

import (
	"fmt"
)

func main() {
	testStrings := []string{
		"главрыба",
		"hello world",
		"🚀🌟 rocket and stars 🌟🚀",
		"Hello, 世界",
		"a",
		"",
		"12345",
	}

	for _, str := range testStrings {
		reversed := reverseString(str)
		fmt.Printf("Original: '%s'\n", str)
		fmt.Printf("Result: '%s'\n\n", reversed)
	}
}

func reverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
