package main

import "fmt"

func groupTemperature(arr []float64) map[int][]float64 {
	res := make(map[int][]float64)

	for _, v := range arr {
		i := int(v) / 10 * 10      // получаем целое число -> получаем частное -> умножаем на 10.
		res[i] = append(res[i], v) // пишем значение в соответствующую группу
	}

	return res
}
func main() {
	var arr = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	result := groupTemperature(arr)
	fmt.Println(result)
}
