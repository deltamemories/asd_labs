package main

import (
	"slices"
)

func Tokenize(str string) []string {
	runes := []rune(str)
	operations := []rune{'+', '-', '/', '*'}
	numbers := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.'}

	// operationsStack := []rune{}
	tokens := []string{}

	for i := 0; i < len(runes); {
		r := runes[i]
		if slices.Contains(operations, r) {
			i++
			tokens = append(tokens, string(r))
		} else if r == '(' {
			i++
			tokens = append(tokens, string(r))
		} else if r == ')' {
			tokens = append(tokens, string(r))
			i++
		} else if slices.Contains(numbers, r) {
			j := i + 1
			for ; j < len(runes); j++ {
				if !slices.Contains(numbers, runes[j]) {
					break
				}
			}
			tokens = append(tokens, string(runes[i:j]))
			i = j

		} else {
			i++
		}
	}

	return tokens
}
