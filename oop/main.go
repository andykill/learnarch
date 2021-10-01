package main

import (
	"bufio"
	"fmt"
	calc2 "homework/oop/calc"
	parse2 "homework/oop/parse"
	stack2 "homework/oop/stack"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

func main() {
	debug := pflag.BoolP("debug", "d", false, "Show stack Pop Push")
	pflag.Parse()

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
		if *debug {
			fmt.Println("Origin string: ", text)
			fmt.Println("Clean string: ", cleanStr)
		}

		parse := parse2.NewParse(cleanStr)
		parse.Debug = *debug

		resultString, err := parse.Str()
		if err != nil {
			fmt.Println("Error: ", err.Error())
			fmt.Println("Result string: ", resultString)
			return
		}
		resultString = strings.Trim(resultString, " ")
		s := strings.Split(resultString, " ")

		stack := stack2.NewStack()
		stack.Debug = *debug
		for _, v := range s {
			if v != "" {
				stack.Push(v)
			}
		}
		calc := calc2.NewCalc(stack)
		calc.Debug = *debug
		result, err := calc.Calculate()
		if err != nil {
			fmt.Println("Error: ", err.Error())
			fmt.Println("Result string: ", resultString)
			fmt.Println("Result Calculate: ", result)
			return
		}
		fmt.Println("Result string: ", resultString)
		fmt.Println("Result Calculate: ", result)
		fmt.Println()
	}
}
