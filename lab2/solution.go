package main

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
)

func tokenize(text string) ([]interface{}, error) {
	numberSymbols := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.'}
	operatorSymbols := []rune{'+', '-', '*', '/', '(', ')'}
	symbolsBeforeUnaryMinus := []rune{'+', '*', '/', '('}
	runes := []rune(text)

	var answer []interface{}
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

func toRpn(tokens []interface{}) ([]interface{}, error) {
	priority := map[string]int{
		"+": 1,
		"-": 1,
		"/": 2,
		"*": 2,
	}
	operationSymbols := []string{"+", "-", "/", "*"}

	stack := []interface{}{}
	queue := []interface{}{}

	for i, t := range tokens {
		switch t := t.(type) {
		case float64:
			queue = append(queue, t)
		case string:
			switch {
			case t == "_":
				if i+1 < len(tokens) {
					n, ok := tokens[i+1].(float64)
					if !ok {
						return nil, errors.New("there should be a number after the minus")
					} else {
						tokens[i+1] = -1 * n
					}
				} else {
					return nil, errors.New("there should be a number after the minus")
				}
			case slices.Contains(operationSymbols, t):
				for len(stack) > 0 {
					op, ok := stack[len(stack)-1].(string)
					if !ok {
						break
					}

					p, ok := priority[op]
					if !ok {
						break
					}

					if p >= priority[t] {
						queue = append(queue, stack[len(stack)-1])
						stack = stack[:len(stack)-1]
					} else {
						break
					}
				}
				stack = append(stack, t)
			case t == "(":
				stack = append(stack, t)
			case t == ")":
				for len(stack) > 0 && stack[len(stack)-1] != "(" {
					queue = append(queue, stack[len(stack)-1])
					stack = stack[:len(stack)-1]
				}
				if len(stack) > 0 {
					stack = stack[:len(stack)-1]
				} else {
					return nil, errors.New("incorrect brackets")
				}
			default:
				return nil, errors.New(fmt.Sprintf("unknown token: %s", t))
			}

		}
	}

	for len(stack) > 0 {
		queue = append(queue, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return queue, nil
}
