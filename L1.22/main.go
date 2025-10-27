package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1 << 30)
	b := big.NewInt(1 << 24)

	fmt.Printf("a = %s\n", a.String())
	fmt.Printf("b = %s\n\n", b.String())

	operations := []struct {
		name string
		op   func(*big.Int, *big.Int, *big.Int) *big.Int
	}{
		{"Plus", (*big.Int).Add},
		{"Minus", (*big.Int).Sub},
		{"Multiply", (*big.Int).Mul},
		{"Divide", (*big.Int).Div},
	}

	for _, operation := range operations {
		result := new(big.Int)
		operation.op(result, a, b)
		fmt.Printf("%s: %s\n", operation.name, result.String())
	}
}
