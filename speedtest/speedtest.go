package speedtest

import (
	"fmt"
	"time"
)

type sortingFunc func([]int) []int

func Speedtest(arr []int, f sortingFunc) ([]int, string) {
	t := time.Now()
	newArr := f(arr)
	return newArr, fmt.Sprintf("Time: %s", time.Since(t))
}
