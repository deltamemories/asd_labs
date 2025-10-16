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

func Sort(arr []int) []int { // fix this
	l := len(arr)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
