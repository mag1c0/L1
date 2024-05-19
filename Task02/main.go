package main

import (
	"fmt"
	"sync"
)

// Функция возведения числа val во вторую степень
func square(val int) int {
	return val * val
}

func main() {
	arr := [...]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22} // Инициалиализируем массив.
	var wg sync.WaitGroup                                   // Инициализируем группу горутин.

	for i, v := range arr { // Бежим циклом по каждому элементу массива.
		wg.Add(1)           // Увеличиваем счетчик.
		go func(i, v int) { // Запускаем анонимную функцию в отдельной горутине.
			defer wg.Done()                                                     // Объявляем отложенную функцию для уменьшения счетчика.
			fmt.Printf("Goroutine %d is finish. Result - %d. \n", i, square(v)) // Выводим результат в stdout.
		}(i, v)
	}
	wg.Wait() // Ожидаем завершения всех горутин из группы wg

	fmt.Println("All jobs done!")
}
