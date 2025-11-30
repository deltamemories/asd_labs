package main

import (
	"asd_labs/speedtest"
	"fmt"
	"math/rand"
)

func main() {
	arr := speedtest.NewRandomIntArray(10, 100)
	fmt.Printf("Array: %v\n", arr)

	quickSortWrapper := func(a []int) []int {
		QuickSort(a)
		return a
	}

	sortedArr, time := speedtest.Speedtest(arr, quickSortWrapper)
	fmt.Printf("Sorted array: %v\n%s\n", sortedArr, time)
}

func QuickSort(arr []int) {
	quickSortRecursive(arr, 0, len(arr)-1)
}

func quickSortRecursive(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSortRecursive(arr, low, pi-1)
		quickSortRecursive(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivotIndex := low + rand.Intn(high-low+1)
pivotValue := arr[pivotIndex]

	arr[pivotIndex], arr[high] = arr[high], arr[pivotIndex]

	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivotValue {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]
	return i
}
