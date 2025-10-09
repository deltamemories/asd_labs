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
	for i := 1; i < len(arr); i++ {
		itemToInsert := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > itemToInsert {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = itemToInsert
	}
	return arr
}
