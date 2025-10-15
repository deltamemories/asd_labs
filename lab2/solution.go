package main

import (
	"errors"
	"fmt"
	"math"
	"slices"
	"strconv"
)

func Calc(text string) (float64, error) {
	runes := []rune(text)
	if len(runes) == 0 {
		return 0.0, errors.New("string is empty")
	}
	if runes[len(runes)-1] != '=' {
		return 0.0, errors.New("string must contain '=' at the end")
	}
	runes = runes[:len(runes)-1]

	t, err := tokenize(runes)
	if err != nil {
		return 0.0, err
	}

	tRpn, err := toRpn(t)
	if err != nil {
		return 0.0, err
	}
	ans, err := calcRpn(tRpn)
	if err != nil {
		return 0.0, err
	}
	return ans, nil
}

func tokenize(runes []rune) ([]interface{}, error) {
	numberSymbols := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.'}
	operatorSymbols := []rune{'+', '-', '*', '/', '(', ')'}
	symbolsBeforeUnaryMinus := []rune{'+', '*', '/', '('}

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

	var stack []interface{}
	var queue []interface{}

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

func add(a, b float64) (float64, error) {
	f := a + b
	if math.IsInf(f, 0) {
		return 0, errors.New("overflow")
	} else {
		return f, nil
	}
}

func sub(a, b float64) (float64, error) {
	f := a - b
	if math.IsInf(f, 0) {
		return 0, errors.New("overflow")
	} else {
		return f, nil
	}
}

func mul(a, b float64) (float64, error) {
	f := a * b
	if math.IsInf(f, 0) {
		return 0, errors.New("overflow")
	} else {
		return f, nil
	}
}

func div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	} else {
		f := a / b
		if math.IsInf(f, 0) {
			return 0, errors.New("overflow")
		} else {
			return f, nil
		}
	}
}

func calcRpn(tokens []interface{}) (float64, error) {
	operationSymbols := []string{"+", "-", "/", "*"}
	var stack []float64

	if len(tokens) == 0 {
		return 0, errors.New("empty tokens")
	}

	for _, t := range tokens {
		switch t := t.(type) {
		case float64:
			stack = append(stack, t)
		case string:
			if slices.Contains(operationSymbols, t) {

				if len(stack) < 2 {
					return 0, errors.New("no enough operands for operator")
				}

				b := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				a := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				var r float64
				var err error
				switch t {
				case "+":
					r, err = add(a, b)
					if err != nil {
						return 0, err
					}
				case "-":
					r, err = sub(a, b)
					if err != nil {
						return 0, err
					}
				case "*":
					r, err = mul(a, b)
					if err != nil {
						return 0, err
					}
				case "/":
					r, err = div(a, b)
					if err != nil {
						return 0, err
					}
				}
				stack = append(stack, r)

			} else {
				return 0, errors.New("unknown token")
			}
		default:
			return 0, errors.New(fmt.Sprintf("unknown token: %s", t))
		}
	}

	if len(stack) > 1 {
		return 0, errors.New("too many operands")
	}

	if len(stack) == 0 {
		return 0, errors.New("empty stack")
	}

	answer := stack[0]
	return answer, nil
}
