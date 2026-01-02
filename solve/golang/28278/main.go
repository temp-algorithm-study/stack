package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	stack := []int{}

	for range n {
		scanner.Scan()
		inputs := strings.Fields(scanner.Text())

		cmd, err := strconv.Atoi(inputs[0])
		if err != nil {
			panic(err)
		}

		switch cmd {
		case 1:
			x, err := strconv.Atoi(inputs[1])
			if err != nil {
				panic(err)
			}
			stack = append(stack, x)
		case 2:
			if len(stack) == 0 {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
		case 3:
			fmt.Fprintln(writer, len(stack))
		case 4:
			if len(stack) == 0 {
				fmt.Fprintln(writer, 1)
			} else {
				fmt.Fprintln(writer, 0)

			}
		case 5:
			if len(stack) == 0 {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, stack[len(stack)-1])
			}
		}

	}
}
