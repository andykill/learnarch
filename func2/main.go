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
		Result(calculate(parseStr(fin())))
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

func parseStr(cleanStr string) (string, *stack.Stack) {
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

func calculate(resultString string, stk *stack.Stack) (string, float64) {
	calc := calc.NewCalc(stk)
	result, err := calc.Calculate()
	if err != nil {
		fmt.Println("Error calculate: ", err.Error())
		fmt.Println("Result string: ", resultString)
	}
	return resultString, result
}

func Result(resultString string, result float64) {
	fmt.Println("Result string: ", resultString)
	fmt.Println("Result Calculate: ", result)
	fmt.Println()
}
