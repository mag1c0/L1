package main

import (
	"fmt"
	"reflect"
)

func getType(v interface{}) {
	fmt.Printf("%T \n", v) // ИЛИ
	fmt.Println(reflect.TypeOf(v))
}
func main() {
	s := []interface{}{1, "", 3.45, false, 'f', struct{}{}, 0b1, []int{2}, [1]int{}}
	for _, v := range s {
		getType(v)
	}
}
