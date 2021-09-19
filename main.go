package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	OpenBracket    = "("
	CloseBracket   = ")"
	PLUS           = "+"
	MINUS          = "-"
	MULTIPLICATION = "*"
	DIVISION       = "/"
)

var priority = map[string]int{
	OpenBracket:    0,
	CloseBracket:   1,
	PLUS:           2,
	MINUS:          2,
	MULTIPLICATION: 3,
	DIVISION:       3,
}

var stack []string
var resultString string
var result int

func main() {
	input := getInput()

	calculate(input)

	output()
}

func output() {
	fmt.Println("Result: ", result)
	fmt.Println("Result string: ", resultString)
}

func calculate(input []string) {
	for _, value := range input {
		calcStr(value)
	}
}

func calcStr(str string) {
	_, err := strconv.Atoi(str)
	if err != nil {
		addToStack(str)
	} else {
		addToResultStr(str)
	}
}

func addToResultStr(str string) {
	resultString += " " + str
}

func addToStack(str string) {
	if str == CloseBracket {
		var strFromStack string
		for {
			strFromStack = pop()
			if strFromStack == OpenBracket {
				return
			}
			addToResultStr(strFromStack)
		}
	}
	stack = append(stack, str)
}

func pop() string {
	n := len(stack) - 1

	str := stack[n]
	stack = stack[:n]
	return str
}

func calc(left int, right int, operator string) int {
	switch operator {
	case PLUS:
		return left + right
	case MINUS:
		return left - right
	case MULTIPLICATION:
		return left * right
	case DIVISION:
		if right == 0 {
			fmt.Println("divide by zero", right)
			os.Exit(1)
		}
		return left / right

	default:
		fmt.Println("Unknown operator ", operator)

	}
	return 0
}

func getInput() []string {
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error input", err)
		os.Exit(1)
	}

	return strings.Split(text, " ")
}
