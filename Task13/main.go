package main

import "fmt"

func main() {
	x, y := 5, 6
	x = x + y
	y = x - y
	x = x - y
	fmt.Println(x, y)

	x, y = 5, 6
	x = x ^ y
	y = y ^ x
	x = x ^ y
	fmt.Println(x, y)

	x, y = 5, 6
	x, y = y, x // Каскадное присваивание
	fmt.Println(x, y)
}
