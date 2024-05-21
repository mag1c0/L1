package main

import (
	"fmt"
	"strconv"
)

func setBit(n int64, i int64, v bool) int64 {
	mask := int64(1) << i // сдвигаем указатель на i. Для примера возмем i = 1. Было 0b0...01 стало 0b0...10

	n &= ^mask // (^mask) - делаем инверсию маски получаем 0b1...01, (n & mask) применяем побитовое И, так мы зануляем значение в индексе i
	// Неважно что будет в исходном бите, он станет 0, т.к 1 & 0 = 0 и 0 & 0 = 0
	//                          *
	// n      11110100001001000000
	// mask   11111111111111111111
	// res    11110100001001000000
	//                          *
	// n      11110100001001000010
	// mask   11111111111111111101
	// res    11110100001001000000

	//Если нужно установить 1, тогда делаем побитовое ИЛИ
	if v {
		n |= mask
	}
	// Неважно что будет в исходном бите, он станет 1, так как (1 | 1 = 1) и (0 | 1 = 1)
	//                          *
	// n      11110100001001000000
	// mask   00000000000000000010
	// res    11110100001001000010
	//                          *
	// n      11110100001001000010
	// mask   00000000000000000010
	// res    11110100001001000010

	return n
}

func main() {
	x := int64(100)
	binary := strconv.FormatInt(x, 2)
	fmt.Println("Исходник              = " + binary)

	x = setBit(x, 1, true)
	binary = strconv.FormatInt(x, 2)
	fmt.Println("Ожидаем 1 в индексе 1 = " + binary)

	x = setBit(x, 1, false)
	binary = strconv.FormatInt(x, 2)
	fmt.Println("Ожидаем 0 в индексе 1 = " + binary)

	x = setBit(x, 3, true)
	binary = strconv.FormatInt(x, 2)
	fmt.Println("Ожидаем 1 в индексе 3 = " + binary)
}
