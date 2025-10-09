package main

import "fmt"

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree", "tree", "treee", "dog"}

	set := make(map[string]bool)

	for _, item := range sequence {
		set[item] = true
	}

	fmt.Print("Result = {")
	first := true
	for key := range set {
		if !first {
			fmt.Print(", ")
		}
		fmt.Printf("%q", key)
		first = false
	}
	fmt.Println("}")
}
