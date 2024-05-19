package main

import (
	"fmt"
	"sync"
	"time"
)

// Функция возведения числа val во вторую степень
func square(val int) int {
	return val * val
}

func main() {
	arr := []int{2, 4, 6, 8, 10} // Инициализируем слайс
	start := time.Now()          // Замерим время выполнения
	firstWay(arr)                // Первая функция возводит исходный массив в квадрат и потом считает его сумму
	duration := time.Since(start)
	fmt.Printf("First way finish - %v\n", duration) // 169.667µs

	arr2 := []int{2, 4, 6, 8, 10} // Инициализируем новый слайс (предыдущий не подходит, т.к элементы в квадрате)
	start = time.Now()            // Замерим время выполнения
	secondWay(arr2)               // Вторая функция работает через объявление канала в который помещается результат горутины и плюсуется
	duration = time.Since(start)
	fmt.Printf("First way finish - %v\n", duration) // 15.958µs
}

func firstWay(arr []int) {
	sum := 0

	var wg sync.WaitGroup
	for i, val := range arr {
		wg.Add(1)
		go func(i, val int) {
			defer wg.Done()
			res := square(val)
			arr[i] = res
		}(i, val)
	}

	wg.Wait()

	for _, val := range arr {
		sum += val
	}

	fmt.Println(sum)
}

func secondWay(arr []int) {
	sum := 0
	var wg sync.WaitGroup
	ch := make(chan int)

	for _, val := range arr {
		wg.Add(1)

		go func(val int, ch chan int) {
			defer wg.Done()
			ch <- square(val)
		}(val, ch)
		sum += <-ch
	}
	wg.Wait()

	fmt.Println(sum)
}
