package main

import "fmt"

func bSearch(arr []int, val int) int {
	if len(arr) < 1 {
		return -1
	}

	l, r := 0, len(arr)
	m := (l + r) / 2

	for l < r {
		if arr[m] == val {
			return m
		} else if arr[m] < val {
			l = m
		} else {
			r = m
		}
		m = (l + r) / 2
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	searchV := 10
	fmt.Println(bSearch(arr, searchV))
	searchV = 2
	fmt.Println(bSearch(arr, searchV))
}
