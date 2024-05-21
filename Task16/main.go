package main

import (
	"fmt"
	"time"
)

func median(arr []int) int {
	if len(arr) < 1 {
		return arr[0]
	}
	i := len(arr) / 2
	if (i % 2) == 0 {
		return arr[i+1]
	} else {
		return arr[i]
	}
}

func partition(arr []int, l, r int) int {
	if r-l <= 1 {
		return 1
	}

	i := l
	j := r - 1
	medArr := make([]int, len(arr[l:r]))
	copy(medArr, arr[l:r])
	med := median(medArr)

	for i <= j {
		for arr[i] < med {
			i++
		}
		for arr[j] > med {
			j--
		}
		if i < j {
			tmp := arr[i]
			arr[i] = arr[j]
			arr[j] = tmp
			i++
			j--
		} else {
			break
		}
	}
	return i
}

func qSortMedian(arr []int, l, r int) {
	if r-l <= 1 {
		return
	}
	m := partition(arr, l, r)
	qSortMedian(arr, l, m)
	qSortMedian(arr, m, r)
}

func qSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var less, greater []int

	for _, v := range arr[1:] {
		if v < pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}
	result := append(qSort(less), pivot)
	result = append(result, qSort(greater)...)

	return result
}
func main() {
	arr := []int{1, 3, 44, 74, 6, 3, 5, 9, 3, 5, 35, 7, 1, 2, 7, 4, 7, 34, 8, 63, 3, 6, 3, 5, 9, 3, 5, 35, 7, 1, 2, 7, 4, 7, 34, 8, 75, 35, 135, 75, 34, 85, 9, 4, 425, 5744, 2}
	start := time.Now()
	qSortMedian(arr, 0, len(arr))
	duration := time.Since(start)
	fmt.Println("duration qsort median: ", duration) // 4.292µs, 13.625µs, 6.209µs, 9.833µs, 12.542µs

	arr2 := []int{1, 3, 44, 74, 6, 3, 5, 9, 3, 5, 35, 7, 1, 2, 7, 4, 7, 34, 8, 63, 3, 6, 3, 5, 9, 3, 5, 35, 7, 1, 2, 7, 4, 7, 34, 8, 75, 35, 135, 75, 34, 85, 9, 4, 425, 5744, 2}
	start = time.Now()
	qSort(arr2)
	duration = time.Since(start)
	fmt.Println("duration qsort: ", duration) // 9.209µs, 18µs, 9.958µs, 13.458µs, 11.25µs
}
