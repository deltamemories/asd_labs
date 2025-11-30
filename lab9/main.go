package main

import (
	"asd_labs/speedtest"
	"fmt"
)

func main() {
	arr := speedtest.NewRandomIntArray(10, 100)
	fmt.Printf("Array: %v\n", arr)

	sortedArr, time := speedtest.Speedtest(arr, HeapSort)
	fmt.Printf("Sorted array: %v\n%s\n", sortedArr, time)
}

func HeapSort(arr []int) []int {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
	return arr
}

func heapify(arr []int, n, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < n && arr[l] > arr[largest] {
		largest = l
	}

	if r < n && arr[r] > arr[largest] {
		largest = r
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}