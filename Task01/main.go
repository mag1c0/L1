package main

import "fmt"

// Структура Human
type Human struct {
	name string
	age  int
}

// Выведем возраст
func (h *Human) printAge() {
	fmt.Println(h.age)
}

// Объявляем структуру Action, с встроенной структурой Human,
// чтобы "наследовать" методы и переменные из родительской структуры
type Action struct {
	Human
}

// Выведем приветствие
func (a *Action) sayHello() {
	fmt.Println("Hello, my name is " + a.name)
}

func main() {
	action := Action{}       // Объявляем структуру Action
	action.name = "John Doe" // Задаем имя (дефолтное значение при инициализации "")
	action.age = 28          // Задаем возраст (дефолтное значение - 0)
	action.printAge()        // 28
	action.sayHello()        // Hello, my name is John Doe
}
