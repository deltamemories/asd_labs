package main

import (
	"asd_labs/speedtest"
	"fmt"
	"math"
)

func main() {
	arr := speedtest.NewRandomIntArray(10, 100)
	fmt.Printf("Array: %v\n", arr)

	sortedArr, time := speedtest.Speedtest(arr, RadixSort)
	fmt.Printf("Sorted array: %v\n%s\n", sortedArr, time)
}

func RadixSort(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	maxDigits := maxLenOfNumbers(arr)
	const base = 10

	for i := 0; i < maxDigits; i++ {
		bins := make([][]int, base)
		for _, x := range arr {
			digit := (x / int(math.Pow(base, float64(i)))) % base
			bins[digit] = append(bins[digit], x)
		}
		arr = []int{}
		for _, bin := range bins {
			arr = append(arr, bin...)
		}
	}

	return arr
}

func countDigits(n int) int {
	if n == 0 {
		return 1
	}
	cnt := 0
	if n < 0 {
		n = -n
	}
	for n > 0 {
		n /= 10
		cnt++
	}
	return cnt
}

func maxLenOfNumbers(arr []int) int {
	if len(arr) == 0 {
		panic("array is empty")
	}

	m := countDigits(arr[0])
	for i := 1; i < len(arr); i++ {
		c := countDigits(arr[i])
		if c > m {
			m = c
		}
	}
	return m
}