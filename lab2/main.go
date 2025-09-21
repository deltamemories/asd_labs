package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	defer fmt.Println("Нажмите enter для выхода")
	defer reader.ReadString('\n')

	fmt.Println("Введите выражение")
	expression, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения консоли")
		return
	}
	expression = strings.TrimSpace(expression)
	fmt.Println("Ваше выражение:", expression)

	fmt.Println(Tokenize(expression))

}
