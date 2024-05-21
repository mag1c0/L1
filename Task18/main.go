package main

import (
	"fmt"
	"sync"
)

type Counter interface { //Объявим интерфейс, с 2-мя методами, т.к в данном примере 2 реализации одного и того же функционала
	Get() interface{}
	Inc()
}

func NewCounterAtomic() Counter {
	return NewAtomicCounter() // Вернем реализацию через Atomic функционал которая удовлетворяет интерфейс
}

func NewCounterMutex() Counter {
	return NewMutexCounter() // Вернем реализацию через sync.Mutex
}

func main() {
	atomCounter := NewCounterAtomic()
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go func(cnt Counter) {
			defer wg.Done()
			cnt.Inc()
		}(atomCounter)
	}
	wg.Wait()
	fmt.Println(atomCounter.Get())

	mutexCounter := NewCounterMutex()
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go func(cnt Counter) {
			defer wg.Done()
			cnt.Inc()
		}(mutexCounter)
	}
	wg.Wait()
	fmt.Println(mutexCounter.Get())
}
