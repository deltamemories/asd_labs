package main

import (
	"errors"
	"slices"
	"strconv"
)

func tokenize(text string) ([]interface{}, error) {
	numberSymbols := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.'}
	operatorSymbols := []rune{'+', '-', '*', '/', '(', ')'}
	symbolsBeforeUnaryMinus := []rune{'+', '*', '/', '('}
	runes := []rune(text)

	answer := []interface{}{}
	i := 0
	for i < len(runes) {
		r := runes[i]
		switch {
		case slices.Contains(operatorSymbols, r):

			if r == '-' {

				if i == 0 || slices.Contains(symbolsBeforeUnaryMinus, runes[i-1]) { // unary minus
					answer = append(answer, "_") // "_" – unary minus sign
				} else {
					answer = append(answer, string(r))
				}

			} else {
				answer = append(answer, string(r))
			}
			i++

		case slices.Contains(numberSymbols, r):
			potentialNumber := []rune{r}
			j := i + 1
			for ; j < len(runes); j++ {
				if slices.Contains(numberSymbols, runes[j]) {
					potentialNumber = append(potentialNumber, runes[j])
				} else {
					break
				}
			}
			n, err := strconv.ParseFloat(string(potentialNumber), 64)
			if err != nil {
				return nil, err
			}
			answer = append(answer, n)
			i = j
		default:
			return nil, errors.New("unknown symbol")
		}
	}
	return answer, nil
}
