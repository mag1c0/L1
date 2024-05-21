package main

import "fmt"

func removeByIndex[T comparable](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}
func main() {
	arrI := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arrS := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	newArrI := removeByIndex(arrI, 3)
	newArrS := removeByIndex(arrS, 5)
	fmt.Println(newArrI)
	fmt.Println(newArrS)
}
