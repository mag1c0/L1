package main

import (
	"fmt"
	"sync"
)

func mutex() {
	mp := make(map[int]int)
	var m sync.Mutex // Инициализируем мьютекс для блокировки доступа
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			m.Lock()      // Блокируем доступ для всех горутин
			mp[val] = val // Записываем данные в мапу
			m.Unlock()    // Снимаем блокировку
		}(i)
	}

	wg.Wait()

	fmt.Println(mp)
}

func channel() {
	mp := make(map[int]int)

	ch := make(chan int) // Инициализируем канал
	var channelWg sync.WaitGroup
	channelWg.Add(1)
	go func() {
		defer channelWg.Done()
		for i := range ch { // Читаем данные из канала и пишем в мапу
			val := i
			mp[val] = val
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			ch <- val // Пишем данные в канал
		}(i)
	}

	wg.Wait()
	close(ch)
	channelWg.Wait()

	fmt.Println(mp)
}

func syncMap() {
	var sm sync.Map // Инициализируем синк мапу (который под капотом использует мьютексы)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			sm.Store(val, val)
		}(i)
	}

	wg.Wait()

	//Для вывода
	m := map[int]int{}
	sm.Range(func(key, value interface{}) bool {
		m[key.(int)] = value.(int)
		return true
	})
	fmt.Println(m)
}

func main() {
	mutex()
	syncMap()
	channel()
}
