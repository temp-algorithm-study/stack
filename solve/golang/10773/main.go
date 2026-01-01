package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sum(slice []int) int {
	result := 0
	for _, v := range slice {
		result += v
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	k, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	stack := []int{}

	for range k {
		scanner.Scan()
		money, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		if money != 0 || len(stack) == 0 {
			stack = append(stack, money)
			continue
		}

		stack = stack[:len(stack)-1]
	}

	fmt.Fprintln(writer, sum(stack))
}
