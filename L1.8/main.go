package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter number: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	num, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		fmt.Println("Error, type correct number")
		return
	}

	fmt.Printf("Binary : %s\n", toBinary(num))

	fmt.Print("Enter bit number: ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	bitPos, err := strconv.Atoi(input)
	if err != nil || bitPos < 0 || bitPos > 63 {
		fmt.Println("Error, bit number must be from 0 to 63")
		return
	}

	fmt.Print("Set bit to 0 or 1: ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var bitValue bool
	if input == "1" {
		bitValue = true
	} else if input == "0" {
		bitValue = false
	} else {
		fmt.Println("Error, type 1 or 0")
		return
	}

	result := setBit(num, bitPos, bitValue)

	fmt.Printf("\nResult:\n")
	fmt.Printf("Start number: %d (%s)\n", num, toBinary(num))
	fmt.Printf("Bit position: %d\n", bitPos)
	fmt.Printf("Set to: %d\n", boolToInt(bitValue))
	fmt.Printf("Result: %d (%s)\n", result, toBinary(result))

}

func setBit(value int64, pos int, bitValue bool) int64 {
	if bitValue {
		return value | (1 << pos)
	} else {
		return value &^ (1 << pos)
	}
}

func toBinary(num int64) string {
	return strconv.FormatInt(num, 2)
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
