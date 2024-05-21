package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseArr(arr []string) []string {
	l, r := 0, len(arr)-1 // Работаем с двумя указателями и меняем их местами
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}

	return arr
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter words: ")
	text, _ := reader.ReadString('\n')

	text = strings.Trim(text, "\n")     // Уберем переносы (т.к читаем с консоли)
	arrText := strings.Split(text, " ") // Разобьем строку на массив

	reverseArr(arrText)                     // разворачиваем элементы
	fmt.Println(strings.Join(arrText, " ")) // склеиваем в строку
}
