package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := make([]int, 10_000)
	for i := range arr {
		arr[i] = int(rand.Int31n(1000000000))
	}

	t := time.Now()
	sortedArr := Sort(arr)
	fmt.Println(sortedArr)
	fmt.Println("Time:", time.Since(t))

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
