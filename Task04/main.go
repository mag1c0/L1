package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Функция отправки данных в канал
func sendData(c chan int, exit chan os.Signal) {
	for {
		select {
		case <-exit: // Если поступила команда, закрываем канал
			close(c)
			return
		default:
			c <- rand.Intn(100)
		}
	}
}

func main() {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, os.Interrupt) // Сигнализируем в канал, если была команда Ctrl + C

	c := make(chan int)   // Инициализируем канал
	var wg sync.WaitGroup // Инициализируем группу

	var n int
	fmt.Print("Введите количество воркеров: ")
	fmt.Fscan(os.Stdin, &n) // Читаем данные из консоли

	for i := 0; i < n; i++ { // Запускаем воркеры
		wg.Add(1)
		go func(mc chan int) {
			defer wg.Done()
			for {
				val, ok := <-mc  // Слушаем канал
				if ok == false { // Если канал закрыт - завершаем работы
					break
				} else { // Если есть данные - выводим в stdout
					fmt.Println(val)
				}
			}
		}(c)
	}

	go sendData(c, exit) // Запускаем отправшик сообщений

	wg.Wait()
}
