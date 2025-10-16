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
	d := len(arr) / 2

	for d >= 1 {
		for i := d; i < len(arr); i++ {
			tmp := arr[i]
			j := i
			for j >= d && arr[j-d] > tmp {
				arr[j] = arr[j-d]
				j -= d
			}
			arr[j] = tmp
		}
		d = d / 2
	}
	return arr
}
