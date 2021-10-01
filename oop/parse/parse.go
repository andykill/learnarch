package parse

import (
	"errors"
	"fmt"
	"homework/oop/stack"
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

type Parse struct {
	str   string
	stack *stack.Stack
	Debug bool
}

func NewParse(str string) *Parse {
	return &Parse{
		str:   str,
		stack: stack.NewStack(),
	}
}
func (parse *Parse) Println(str ...interface{}) {
	if parse.Debug {
		fmt.Println(str)
	}
}

func (parse *Parse) Str() (string, error) {
	resultString := strings.Builder{}

	for idx := 0; idx < len([]rune(parse.str)); idx++ {
		curr := string(parse.str[idx])

		parse.Println("Current: ", curr)
		switch {
		case curr == OpenBracket:
			parse.stack.Push(OpenBracket)
		case curr == CloseBracket:
			for {
				strFromStack := parse.stack.Pop()

				if strFromStack == OpenBracket {
					break
				}
				if parse.stack.IsEmpty() {
					return "", errors.New("Not found OpenBracket")
				}
				parse.Println("WriteString " + strFromStack)
				resultString.WriteString(strFromStack)
				resultString.WriteString(" ")
			}
		case isInt(curr):
			parse.Println("WriteString " + curr)
			resultString.WriteString(curr)
			for {
				n, ok := next(parse.str, idx)
				if ok && (isInt(n) || n == ".") {
					parse.Println("WriteString " + n)
					resultString.WriteString(n)
					idx++
				} else {
					break
				}
			}
		case isOperator(curr):
			if isUnaryOperation(parse.str, idx) {
				parse.Println("WriteString " + curr)
				resultString.WriteString(curr)
				for {
					if n, ok := next(parse.str, idx); ok && isInt(n) {
						parse.Println("WriteString " + n)
						resultString.WriteString(n)
						resultString.WriteString(" ")
						idx++
					} else {
						break
					}
				}
			} else {
				operatorFromStack := parse.stack.Pop()
				if isPriority(curr, operatorFromStack) {
					parse.stack.Push(operatorFromStack)
					parse.stack.Push(curr)
				} else {
					parse.Println("WriteString " + operatorFromStack)
					resultString.WriteString(operatorFromStack)
					parse.stack.Push(curr)
				}
			}
		default:
			return "", errors.New("Found element not operator and not digital")
		}

		resultString.WriteString(" ")
		parse.Println("Stack strings out: ", parse.stack.Strings())
		parse.Println("Result string out: ", resultString.String())
		parse.Println()
	}

	for {
		if parse.stack.IsEmpty() {
			break
		}
		resultString.WriteString(parse.stack.Pop())
		resultString.WriteString(" ")
	}
	return resultString.String(), nil
}

func isPriority(curr string, operator string) bool {

	if curr == "" {
		return true
	}

	if operator == "" {
		return false
	}

	p1, ok := priority[curr]
	if !ok {
		panic("Priority not found" + curr)
	}
	p2, ok := priority[operator]
	if !ok {
		panic("Priority not found" + curr + " " + operator)
	}
	return p1 > p2
}

func pop(stack []string) ([]string, string) {
	n := len(stack) - 1
	if n < 0 {
		return stack, ""
	}
	str := stack[n]
	stack = stack[:n]

	return stack, str
}

func next(str string, idx int) (string, bool) {
	lenStr := len(str)
	if lenStr > idx+1 {
		return str[idx+1 : idx+2], true
	}
	return "", false
}
func prev(str string, idx int) (string, bool) {
	if idx-1 >= 0 {
		return str[idx-1 : idx], true
	}
	return "", false
}

func isUnaryOperation(str string, idx int) bool {
	curr := string(str[idx])
	if !isUnaryOperator(curr) {
		return false
	}

	if p, ok := prev(str, idx); !ok {
		if n, ok := next(str, idx); ok && isInt(n) {
			return true
		}
	} else if isOperator(p) {
		return true
	}
	return false
}

func isUnaryOperator(operator string) bool {
	switch operator {
	case PLUS:
		return true
	case MINUS:
		return true
	}
	return false
}

func isOperator(operator string) bool {
	switch operator {
	case PLUS:
		return true
	case MINUS:
		return true
	case MULTIPLICATION:
		return true
	case DIVISION:
		return true
	}
	return false
}

func isInt(v string) bool {
	if _, err := strconv.ParseFloat(v, 64); err == nil {
		return true
	}
	return false
}
