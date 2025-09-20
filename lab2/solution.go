package main

import (
	"fmt"
	"slices"
)

func Calc(str string) (float64, error) {
	runes := []rune(str)
	operations := []rune{'+', '-', '/', '*'}
	numbers := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.'}

	// operationsStack := []rune{}
	numbersStack := []string{}

	for i := 0; i < len(runes); {
		r := runes[i]
		if slices.Contains(operations, r) {
			i++
		} else if r == '(' {
			i++
		} else if r == ')' {
			i++
		} else {
			j := i + 1
			for ; j < len(runes); j++ {
				if !slices.Contains(numbers, runes[j]) {
					break
				}
			}

			fmt.Println(i, j, string(runes[i:j]))
			numbersStack = append(numbersStack, string(runes[i:j]))
			i = j

		}
	}

	fmt.Println("numbersStack")
	fmt.Println(numbersStack)

	return 0.0, nil
}
