package main

import (
	"errors"
	"fmt"
	"slices"
)

func Tokenize(str string) ([]string, error) {
	runes := []rune(str)
	operations := []rune{'+', '-', '/', '*'}
	numbers := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.'}

	tokens := []string{}

	if len(runes) != 0 && runes[len(runes)-1] != '=' {
		return []string{}, errors.New("the expression must end with '=' sign")
	}

	if len(runes) != 0 {
		runes = runes[:len(runes)-1]
	}

	for i := 0; i < len(runes); {
		r := runes[i]
		if slices.Contains(operations, r) {
			tokens = append(tokens, string(r))
			i++
		} else if r == '(' {
			tokens = append(tokens, string(r))
			i++
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
		} else if r == ' ' {
			i++
		} else {
			return []string{}, errors.New("unknown char in expression")
		}
	}

	return tokens, nil
}

// func Calc(tokens []string) (float64, error) {
// 	// return 0.0, errors.New("divide by zero")

// 	operationsStack := []string{}
// 	stack := []string{}

// 	for _, t := range tokens {

// 	}
// }

func ToRpn(tokens []string) []string {
	operations := []string{"+", "-", "/", "*"}
	numbers := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "."}

	prec := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}

	stack := []string{}
	queue := []string{}

	for _, t := range tokens {
		fmt.Println("---------")
		fmt.Println(t, stack, queue)
		fmt.Println("---------")

		if slices.Contains(numbers, t) {
			queue = append(queue, t)
		}

		if slices.Contains(operations, t) {
			fmt.Println(prec[stack[len(stack)-1]])
			// len(stack) > 0 && prec[stack[len(stack)-1]] >= prec[t]
			for len(stack) > 0 {
				p, ok := prec[stack[len(stack)-1]]

				if !ok {
					break
				}

				if p >= prec[t] {
					queue = append(queue, stack[len(stack)-1])
					stack = stack[:len(stack)-1]
				} else {
					break
				}
			}
			stack = append(stack, t)
		}

		if t == "(" {
			stack = append(stack, t)
		}
		if t == ")" {
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				queue = append(queue, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		}

	}

	for len(stack) > 0 {
		queue = append(queue, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return queue
}

func sum(a float64, b float64) float64 {
	return a + b
}

func sub(a float64, b float64) float64 {
	return a - b
}

func mul(a float64, b float64) float64 {
	return a * b
}

func div(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0.0, errors.New("divide by zero")
	} else {
		return a / b, nil
	}
}
