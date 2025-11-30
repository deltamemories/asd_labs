package main

import (
	"asd_labs/speedtest"
	"fmt"
)

func main() {
	arr := speedtest.NewRandomIntArray(10, 100)
	fmt.Printf("Array: %v\n", arr)

	sortedArr, time := speedtest.Speedtest(arr, MergeSort)
	fmt.Printf("Sorted array: %v\n%s\n", sortedArr, time)
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}