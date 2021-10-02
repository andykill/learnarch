package main

import (
	"bufio"
	"fmt"
	"homework/func/calc"
	"homework/func/parse"
	"homework/func/stack"
	"os"
	"strings"
)

func main() {
	//cleanStr = "(666+10-4)/(1+1*2)+1"
	//cleanStr = "3+3*3"

	for {
		Result(calculate, fin)
	}
}

func fin() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Write a line that needs to be calculated, or !q to exit")
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	replacer := strings.NewReplacer(",", ".", " ", "", "\n", "", "\t", "")
	cleanStr := replacer.Replace(text)
	if cleanStr == "!q" {
		fmt.Println("Good bye")
		os.Exit(1)
	}
	return cleanStr
}

func parseStr(f func() string) (string, *stack.Stack) {
	cleanStr := f()

	parse := parse.NewParse(cleanStr)
	stk := stack.NewStack()

	resultString, err := parse.Str()
	if err != nil {
		fmt.Println("Error parseStr: ", err.Error())
		fmt.Println("Result string: ", resultString)
		return "", stk
	}

	resultString = strings.Trim(resultString, " ")
	s := strings.Split(resultString, " ")
	for _, v := range s {
		if v != "" {
			stk.Push(v)
		}
	}

	return resultString, stk
}

func calculate(parseStr func(func() string) (string, *stack.Stack), fin func() string) (string, float64) {
	resultString, stk := parseStr(fin)
	calc := calc.NewCalc(stk)
	result, err := calc.Calculate()
	if err != nil {
		fmt.Println("Error calculate: ", err.Error())
		fmt.Println("Result string: ", resultString)
	}
	return resultString, result
}

func Result(calculate func(func(func() string) (string, *stack.Stack), func() string) (string, float64), fin func() string) {

	resultString, result := calculate(parseStr, fin)
	fmt.Println("Result string: ", resultString)
	fmt.Println("Result Calculate: ", result)
	fmt.Println()
}
