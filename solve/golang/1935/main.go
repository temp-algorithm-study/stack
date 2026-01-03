package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	// 1. n 입력 받기
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	// 2. 후위 표기식 입력받기
	scanner.Scan()
	postfix := []rune(scanner.Text())
	postfixMap := make(map[rune]float64)

	// 3. 후위 표기식에 값 대입
	for i := 0; i < n; i++ {
		scanner.Scan()
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		postfixMap[rune(65+i)] = float64(val)
	}

	stack := []float64{}
	for _, ch := range postfix {
		if ch >= 'A' && ch <= 'Z' {
			stack = append(stack, postfixMap[ch])
			continue
		}
		operand2 := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		operand1 := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		switch ch {
		case '+':
			stack = append(stack, operand1+operand2)
		case '-':
			stack = append(stack, operand1-operand2)
		case '*':
			stack = append(stack, operand1*operand2)
		case '/':
			stack = append(stack, operand1/operand2)
		}
	}

	fmt.Fprintf(writer, "%.2f\n", stack[0])
}
