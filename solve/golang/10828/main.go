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

	for i := 0; i < n; i++ {
		scanner.Scan()
		cmd := strings.Fields(scanner.Text())

		switch cmd[0] {
		case "push":
			val, err := strconv.Atoi(cmd[1])
			if err != nil {
				panic(err)
			}
			stack = append(stack, val)
		case "pop":
			if len(stack) == 0 {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
		case "size":
			fmt.Fprintln(writer, len(stack))
		case "empty":
			if len(stack) == 0 {
				fmt.Fprintln(writer, 1)

			} else {
				fmt.Fprintln(writer, 0)
			}
		case "top":
			if len(stack) == 0 {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, stack[len(stack)-1])
			}
		}
	}
}
