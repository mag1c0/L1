package main

import (
	"fmt"
	"strings"
)

func checkUnique(s string) bool {
	store := make(map[rune]struct{})
	s = strings.ToLower(s)

	for _, v := range s {
		if _, ok := store[v]; ok {
			return false
		}
		store[v] = struct{}{}
	}

	return true
}

func main() {
	s := "abcd"
	fmt.Println(checkUnique(s)) // true
	s = "abCdefAaf"
	fmt.Println(checkUnique(s)) // false
	s = "aabcd"
	fmt.Println(checkUnique(s)) // false

	s = "abcdeyrlv"
	fmt.Println(checkUnique(s)) // true
	s = "Vasfv"
	fmt.Println(checkUnique(s)) // false
	s = "123456"
	fmt.Println(checkUnique(s)) // true
	s = "12347631"
	fmt.Println(checkUnique(s)) // false
}
