package speedtest

import (
	"fmt"
	"math/rand"
	"time"
)

type sortingFunc func([]int) []int

func Speedtest(arr []int, f sortingFunc) ([]int, string) {
	t := time.Now()
	newArr := f(arr)
	return newArr, fmt.Sprintf("Time: %s", time.Since(t))
}

func NewRandomIntArray(l int, maxVal int) []int {
	arr := make([]int, l)
	for i := range arr {
		arr[i] = int(rand.Intn(maxVal))
	}
	return arr
}
