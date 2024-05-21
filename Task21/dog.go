package main

import "fmt"

type Dog struct{}

func (d *Dog) Woof() {
	fmt.Println("Wof")
}
