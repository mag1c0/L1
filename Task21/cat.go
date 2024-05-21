package main

import "fmt"

type Cat struct{}

func (c *Cat) Meow() {
	fmt.Println("Meow")
}
