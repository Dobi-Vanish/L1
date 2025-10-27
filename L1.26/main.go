package main

import (
    "fmt"
    "strings"
)

func hasAllUniqueChars(str string) bool {
    str = strings.ToLower(str)
    
    seen := make(map[rune]bool)
    
    for _, char := range str {
        if seen[char] {
            return false
        }
        seen[char] = true
    }
    
    return true
}

func main() {
    testCases := []string{
        "abcd",
        "abCdefAaf",
        "aabcd",
        "abcde",
        "Hello",
        "WORLD",
        "",
        "a",
    }
    
    for _, test := range testCases {
        fmt.Printf("'%s' -> %t\n", test, hasAllUniqueChars(test))
    }
}