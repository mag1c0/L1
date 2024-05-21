package main

import (
	"fmt"
	"time"
)

func sender(cOut chan int, arr []int) {
	timeout := time.After(15 * time.Second) // Ограничим выполнение горутины таймаутом. т.к неизвестно какой будет массив
	defer close(cOut)
	defer fmt.Println("sender goroutine exited")

	for _, v := range arr { // Когда все элементы отправлены(либо таймаут) - закрываем канал и выходим
		select {
		case <-timeout:
			return
		default:
			cOut <- v
		}
	}
}

func reader(cIn chan int, cOut chan int) {
	defer close(cOut)
	defer fmt.Println("reader goroutine exited")

	for {
		val, ok := <-cIn // Слушаем канал
		if ok == false { // Если канал закрыт - закрываем cOut канал (cIn уже будет закрыт) - выходим
			return
		} else { // Если есть данные - передаем в канал
			cOut <- val * 2
		}
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}

	cIn := make(chan int)
	cOut := make(chan int)

	go sender(cIn, arr)
	go reader(cIn, cOut)

	for v := range cOut { // Читаем все данные до закрытия из горутины(reader).
		fmt.Println(v)
	}
}
