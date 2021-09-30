package main

import (
	"bufio"
	"errors"
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Write a line that needs to be calculated, or !q to exit")
		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		replacer := strings.NewReplacer(",", ".", " ", "", "\n", "", "\t", "")
		cleanStr := replacer.Replace(text)
		if cleanStr == "!q" {
			fmt.Println("Good bye")
			return
		}
		//cleanStr = "(666+10-4)/(1+1*2)+1"
		//cleanStr = "3+3*3"
		fmt.Println("Origin string: ", text)
		fmt.Println("Clean string: ", cleanStr)

		resultString, err := parseStr(cleanStr)
		if err != nil {
			panic(err)
			return
		}
		resultString = strings.Trim(resultString, " ")
		s := strings.Split(resultString, " ")

		var arrStr []string
		for _, v := range s {
			if v != "" {
				arrStr = append(arrStr, v)
			}
		}

		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}

		result, err := calculate(arrStr)
		if err != nil {
			panic(err)
			return
		}
		fmt.Println("Result string: ", resultString)
		fmt.Println("Result Calculate: ", result)
		fmt.Println()
	}
}

func calculate(arrString []string) (float64, error) {

	var err error
	if arrString == nil {
		return 0, nil
	}
	arrString, operator := pop(arrString)

	arrString, right := pop(arrString)
	var rightInt float64
	if !isInt(right) {
		arrString = append(arrString, right)
		rightInt, err = calculate(arrString)
		if err != nil {
			panic("Error calculate " + err.Error())
		}
		arrString, _ = pop(arrString)
	} else {
		rightInt, _ = strconv.ParseFloat(right, 64)
	}

	arrString, left := pop(arrString)
	var leftInt float64
	if !isInt(left) {
		arrString = append(arrString, left)
		leftInt, err = calculate(arrString)
		arrString, _ = pop(arrString)
		if err != nil {
			panic("Error calculate " + err.Error())
		}
	} else {
		leftInt, _ = strconv.ParseFloat(left, 64)
	}
	//fmt.Println(leftInt)
	//fmt.Println(rightInt)
	//fmt.Println(operator)

	return calc(leftInt, rightInt, operator)
}

// (666+-10*4)/(1+1*2)+1
// 6 10 + 4 - 1 1 2 * + / 1 +

// Вывод: если текущий знак является операцией, а последний знак из стека операций имеет приоритет ниже или равный,
//то последний знак из стека уходит в массив выхода,
//а текущий добавляется в стек. В ином случае все добавляется как обычно.
//Проще говоря: операции — это кандидаты на добавление в массив выхода,
//но чтобы им туда попасть, нужен знак с меньшим, или с таким же приоритетом на входе.
func parseStr(str string) (string, error) {
	var stack []string
	resultString := strings.Builder{}

	strRune := []rune(str)
	lenStr := len(strRune)

	for idx := 0; idx < lenStr; idx++ {
		curr := string(str[idx])

		//fmt.Println("Current: ", curr)
		//fmt.Println("Stack: ", stack)
		//fmt.Println("Result string: ", resultString.String())
		if curr == OpenBracket {
			stack = append(stack, OpenBracket)
			continue
		}

		if curr == CloseBracket {
			for {
				var strFromStack string
				stack, strFromStack = pop(stack)

				if strFromStack == OpenBracket {
					//stack, _ = pop(stack)
					break
				}
				resultString.WriteString(strFromStack)
				resultString.WriteString(" ")
			}
			continue
		}
		if isInt(curr) {
			resultString.WriteString(curr)
			for {
				n, ok := next(str, idx)
				if ok && (isInt(n) || n == ".") {
					resultString.WriteString(n)
					idx++
				} else {
					break
				}
			}
		} else if !isOperator(curr) {
			panic("Found element not operator and not digital")
		} else {

			if p, ok := prev(str, idx); ok && isOperator(p) && isUnaryOperator(curr) {
				resultString.WriteString(curr)
				for {
					if next, ok := next(str, idx); ok && isInt(next) {
						resultString.WriteString(next)
						idx++
					} else {
						break
					}
				}
			} else {
				_, operatorInStack := pop(stack)
				if isPriority(curr, operatorInStack) {
					stack = append(stack, curr)
				} else {
					resultString.WriteString(operatorInStack)
					stack, _ = pop(stack)
					stack = append(stack, curr)
				}
			}
		}

		resultString.WriteString(" ")
		//fmt.Println("Stack out : ", stack)
		//fmt.Println("Result string out: ", resultString.String())
		//fmt.Println()
	}
	for _ = range stack {
		var v string
		stack, v = pop(stack)
		resultString.WriteString(v)
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

func calc(left float64, right float64, operator string) (float64, error) {
	switch operator {
	case PLUS:
		return float64(left + right), nil
	case MINUS:
		return float64(left - right), nil
	case MULTIPLICATION:
		return float64(left * right), nil
	case DIVISION:
		if right == 0 {
			return 0, errors.New("divide by zero:")
		}
		return float64(left / right), nil
	}
	return 0, errors.New("Unknown operator: " + operator)
}
