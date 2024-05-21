package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, ok := new(big.Int).SetString("240000000000000000", 10)
	if !ok {
		fmt.Println("Error")
		return
	}

	b, ok := new(big.Int).SetString("240000000000000000", 10)
	if !ok {
		fmt.Println("Error")
		return
	}

	c := new(big.Int)
	fmt.Println(c.Add(a, b))
	fmt.Println(c.Sub(a, b))
	fmt.Println(c.Div(a, b))
	fmt.Println(c.Mul(a, b))
}
