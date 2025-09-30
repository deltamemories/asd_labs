package main

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
)

func Calc(str string) (float64, error) {
	tokens, err := tokenize(str)
	if err != nil {
		return 0.0, err
	}
	fmt.Println("TOKENS:", tokens)

	rpn, err := toRpn(tokens)
	if err != nil {
		return 0.0, err
	}
	fmt.Println("RPN:", rpn)
	result, err := calcRpn(rpn)
	if err != nil {
		return 0.0, err
	}
	return result, nil
}

func tokenize(str string) ([]string, error) {
	runes := []rune(str)
	operations := []rune{'+', '-', '/', '*'}
	numbers := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.'}

	tokens := []string{}

	if len(runes) == 0 {
		return []string{}, errors.New("empty string")
	}

	if len(runes) != 0 && runes[len(runes)-1] != '=' {
		return []string{}, errors.New("the expression must end with '=' sign")
	}

	if len(runes) != 0 {
		runes = runes[:len(runes)-1]
	}

	for i := 0; i < len(runes); {
		r := runes[i]
		if slices.Contains(operations, r) {
			if r == '-' && (i == 0 || (i > 0 && !slices.Contains(numbers, runes[i-1]))) {
				tokens = append(tokens, "(")
				tokens = append(tokens, "0")
				tokens = append(tokens, string(r))
				tokens = append(tokens, ")")
				i++
			} else {
				tokens = append(tokens, string(r))
				i++
			}
		} else if r == '(' {
			if i > 0 && (slices.Contains(numbers, runes[i-1]) || runes[i-1] == ')') {
				tokens = append(tokens, "*")
				tokens = append(tokens, string(r))
				i++
			} else {
				tokens = append(tokens, string(r))
				i++
			}
		} else if r == ')' {
			if i+1 < len(runes) && (slices.Contains(numbers, runes[i+1]) || runes[i+1] == '(') {
				tokens = append(tokens, string(r))
				tokens = append(tokens, "*")
				i++
			} else {
				tokens = append(tokens, string(r))
				i++
			}
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

func calcRpn(tokens []string) (float64, error) {
	operations := []string{"+", "-", "/", "*"}

	stack := []string{}

	if len(tokens) == 0 {
		return 0.0, nil
	}

	for _, t := range tokens {
		_, err := strconv.ParseFloat(t, 64)
		if err == nil {
			stack = append(stack, t)
		} else if slices.Contains(operations, t) {
			if len(stack) < 2 {
				return 0.0, errors.New("no enought operands for operator")
			}

			b, err := strconv.ParseFloat(stack[len(stack)-1], 64)
			if err != nil {
				return 0.0, err
			}
			stack = stack[:len(stack)-1]

			a, err := strconv.ParseFloat(stack[len(stack)-1], 64)
			if err != nil {
				return 0.0, err
			}
			stack = stack[:len(stack)-1]

			r := 0.0

			switch t {
			case "+":
				r = sum(a, b)
			case "-":
				r = sub(a, b)
			case "*":
				r = mul(a, b)
			case "/":
				r, err = div(a, b)
				if err != nil {
					return 0.0, err
				}
			}
			stack = append(stack, strconv.FormatFloat(r, 'f', -1, 64))
		} else {
			return 0.0, errors.New("unknown token")
		}
	}
	if len(stack) > 1 {
		return 0.0, errors.New("too much operands")
	}
	answer, err := strconv.ParseFloat(stack[0], 64)
	if err != nil {
		return 0.0, err
	}
	return answer, nil
}

func toRpn(tokens []string) ([]string, error) {
	operations := []string{"+", "-", "/", "*"}

	prec := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}

	stack := []string{}
	queue := []string{}

	for _, t := range tokens {
		_, err := strconv.ParseFloat(t, 64)
		if err == nil {
			queue = append(queue, t)
		} else if slices.Contains(operations, t) {
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
		} else if t == "(" {
			stack = append(stack, t)
		} else if t == ")" {
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				queue = append(queue, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				return []string{}, errors.New("incorrect brackets")
			}
		} else {
			return []string{}, errors.New("inknown char")
		}

	}

	for len(stack) > 0 {
		queue = append(queue, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return queue, nil
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
