package main

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100] // При такой инициализации, подстрока justString ссылается на часть основной строки "v".
	// После завершении функции, большая область памяти, выделенная на переменную 'v' не очистится, пока justString не перестанет ссылаться на неё.

	// Решение - скопировать нужную подстроку в новую переменную
	byteSlice := make([]byte, 100)
	copy(byteSlice, v[:100])
	justString = string(byteSlice)
}

func main() {
	someFunc()
}
