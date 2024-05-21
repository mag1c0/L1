package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	m := make(map[string]struct{})

	for _, v := range arr {
		m[v] = struct{}{}
	}

	fmt.Println(m)
}
