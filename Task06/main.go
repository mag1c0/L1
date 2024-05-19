package main

import (
	"context"
	"fmt"
	"time"
)

func stopWithChan() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop: // Если поступила команда - заканчиваем
				close(stop)
				fmt.Println("stop goroutine with channel")
				return
			default:
				fmt.Println("running")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	stop <- true // Отправляем любой сигнал в канал, чтобы остановить выполнение горутины
	time.Sleep(5 * time.Second)
}

func stopWithContextCancel() {
	ctx, cancel := context.WithCancel(context.Background()) // Объявляем контекс

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stop goroutine with context.Done")
				return
			default:
				fmt.Println("running")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(5 * time.Second)
}

func stopWithContextTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Объявляем контекс
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stop goroutine with context timeout")
				return
			default:
				fmt.Println("running")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
}

func stopWithTimeout() {
	timeout := time.After(5 * time.Second) // Объявляем канал с таймером

	go func() {
		for {
			select {
			case <-timeout: // Время вышло, завершаем горутину
				fmt.Println("stop goroutine with timeout")
				return
			default:
				fmt.Println("running")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
}

func main() {
	stopWithChan()
	stopWithContextCancel()  // Два одинаковых метода, с небольшими изменениями но тоже имеет место быть
	stopWithContextTimeout() // Два одинаковых метода, с небольшими изменениями но тоже имеет место быть
	stopWithTimeout()
}
