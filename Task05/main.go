package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	N = 5
)

func publisher(ctx context.Context, wg *sync.WaitGroup, c chan int) {
	ctx, cancel := context.WithTimeout(ctx, N*time.Second)
	defer cancel()
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			close(c)
			return
		default:
			c <- rand.Intn(100)
			time.Sleep(time.Second * 1)
		}
	}
}

func subscriber(ctx context.Context, wg *sync.WaitGroup, c chan int) {
	ctx, cancel := context.WithTimeout(ctx, N*time.Second)
	defer cancel()
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			val, ok := <-c   // Слушаем канал
			if ok == false { // Если канал закрыт - завершаем работу
				return
			} else { // Если есть данные - выводим в stdout
				fmt.Println(val)
			}
		}
	}
}

func main() {
	ctx := context.Background()
	var wg sync.WaitGroup
	c := make(chan int) // Инициализация канала

	wg.Add(2)
	go subscriber(ctx, &wg, c) // Подписываемся на канал
	go publisher(ctx, &wg, c)  // Посылаем данные в канал

	wg.Wait()
}
