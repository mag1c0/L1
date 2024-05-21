package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverse(s string) string {
	var result = []rune(s)
	l, r := 0, len(result)-1 // Работаем с двумя указателями и меняем их местами
	for l < r {
		result[l], result[r] = result[r], result[l]
		l++
		r--
	}

	return string(result)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter word: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	fmt.Println(reverse(text))
}
