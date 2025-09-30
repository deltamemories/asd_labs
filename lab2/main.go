package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	defer reader.ReadString('\n')
	defer fmt.Println("Нажмите enter для выхода")

	fmt.Println("math.MaxFloat64", math.MaxFloat64)

	fmt.Println("Введите выражение")
	expression, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения консоли")
		return
	}
	expression = strings.TrimSpace(expression)
	fmt.Println("Ваше выражение:", expression)

	ans, err := Calc(expression)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:")
		fmt.Println(ans)
	}

}
