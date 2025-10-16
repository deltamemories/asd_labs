package main

import (
	"asd_labs/speedtest"
	"fmt"
)

func main() {
	arr := speedtest.NewRandomIntArray(10, 100)
	fmt.Printf("Array: %v\n", arr)

	sortedArr, time := speedtest.Speedtest(arr, Sort)
	fmt.Printf("Sorted array: %v\n%s\n", sortedArr, time)
}

func Sort(arr []int) []int {
	fac := 1.247
	l := len(arr)
	step := l
	flag := false

	for step > 1 || flag {
		if step > 1 {
			step = int(float64(step) / fac)
		}
		flag = false

		for i := 0; i+step < l; i += step {
			if arr[i] > arr[i+step] {
				arr[i], arr[i+step] = arr[i+step], arr[i]
				flag = true
			}
		}

	}
	return arr
}
