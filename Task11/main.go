package main

import "fmt"

func setIntersection(a1, a2 []int) []int {
	var result []int

	for _, v := range a1 { // Пробегаем по первому массиву с внутренним циклом который проверяет пересечение с элементами второго массива. O(N²)
		for _, j := range a2 {
			if v == j {
				result = append(result, j)
			}
		}
	}

	return result
}
func main() {
	a1 := []int{1, 3, 5, 2, 9, 6}
	a2 := []int{2, 6, 1, 4, 10, 9}

	fmt.Println(setIntersection(a1, a2))
}
