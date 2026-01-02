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
	// n, err := strconv.Atoi(scanner.Text())
	// if err != nil {
	// 	panic(err)
	// }

	stack := []int{}

	scanner.Scan()
	currentStudentList := strings.Fields(scanner.Text())
	expacted := 1
	for _, student := range currentStudentList {
		currentStudent, err := strconv.Atoi(student)
		if err != nil {
			panic(err)
		}

		stack = append(stack, currentStudent)

		for len(stack) > 0 && stack[len(stack)-1] == expacted {
			stack = stack[:len(stack)-1]
			expacted++
		}
	}

	if len(stack) == 0 {
		fmt.Fprintln(writer, "Nice")
	} else {
		fmt.Fprintln(writer, "Sad")
	}

}
