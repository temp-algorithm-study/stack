package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxCapacity = 600000

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 1. 초기 문장 입력
	scanner.Scan()
	leftStack := []rune(scanner.Text())
	rightStack := []rune{}

	// 2. 명령어 수 입력
	scanner.Scan()
	m, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	for i := 0; i < m; i++ {
		// 3. 명령 입력받기
		scanner.Scan()
		cmdLine := []rune(strings.ReplaceAll(scanner.Text(), " ", ""))
		switch cmdLine[0] {
		case 'L':
			if len(leftStack) == 0 {
				continue
			}

			ch := leftStack[len(leftStack)-1]
			leftStack = leftStack[:len(leftStack)-1]
			rightStack = append(rightStack, ch)

		case 'D':
			if len(rightStack) == 0 {
				continue
			}

			ch := rightStack[len(rightStack)-1]
			rightStack = rightStack[:len(rightStack)-1]
			leftStack = append(leftStack, ch)

		case 'B':
			if len(leftStack) == 0 {
				continue
			}
			leftStack = leftStack[:len(leftStack)-1]

		case 'P':
			leftStack = append(leftStack, cmdLine[1])

		}
	}

	fmt.Fprintf(writer, "%s", string(leftStack))
	for i := len(rightStack) - 1; i >= 0; i-- {
		writer.WriteRune(rightStack[i])
	}
	fmt.Fprint(writer, "\n")
}
