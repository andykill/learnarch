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
		Result(func(fin func() string) func() (string, float64) {
			cleanStr := fin()

			prs := parse.NewParse(cleanStr)
			stk := stack.NewStack()

			resultString, err := prs.Str()
			if err != nil {
				fmt.Println("Error parseStr: ", err.Error())
				fmt.Println("Result string: ", resultString)
				os.Exit(1)
			}

			resultString = strings.Trim(resultString, " ")
			s := strings.Split(resultString, " ")
			for _, v := range s {
				if v != "" {
					stk.Push(v)
				}
			}

			return func() (string, float64) {
				clc := calc.NewCalc(stk)
				result, err := clc.Calculate()
				if err != nil {
					fmt.Println("Error calculate: ", err.Error())
					fmt.Println("Result string: ", resultString)
				}
				return resultString, result
			}
		},
			func() string {
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
			},
		)
	}
}

func Result(
	parseStr func(f func() string) func() (string, float64),
	fin func() string) {

	calculate := parseStr(fin)
	resultString, result := calculate()
	fmt.Println("Result string: ", resultString)
	fmt.Println("Result Calculate: ", result)
	fmt.Println()
}
