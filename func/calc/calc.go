package calc

import (
	"errors"
	"fmt"
	"homework/func/stack"
	"strconv"
)

type Calculator struct {
	stack *stack.Stack
	Debug bool
}

func NewCalc(stk *stack.Stack) *Calculator {
	return &Calculator{
		stack: stk,
	}
}

func (calc *Calculator) Calculate() (float64, error) {

	var err error
	if calc.stack.IsEmpty() {
		return 0, nil
	}
	operator := calc.stack.Pop()

	right := calc.stack.Pop()
	var rightInt float64
	if !isNumeric(right) {
		calc.stack.Push(right)

		clone := calc.stack.Clone()
		newCalc := NewCalc(&clone)
		rightInt, err = newCalc.Calculate()
		if err != nil {
			return 0, errors.New("Error calculate " + err.Error())
		}
		calc.stack.Pop()
	} else {
		rightInt, _ = strconv.ParseFloat(right, 64)
	}

	left := calc.stack.Pop()
	var leftInt float64
	if !isNumeric(left) {
		calc.stack.Push(left)
		clone := calc.stack.Clone()
		newCalc := NewCalc(&clone)
		leftInt, err = newCalc.Calculate()
		if err != nil {
			return 0, errors.New("Error calculate " + err.Error())
		}
		calc.stack.Pop()
	} else {
		leftInt, _ = strconv.ParseFloat(left, 64)
	}

	if calc.Debug {
		fmt.Println("leftInt: ", leftInt)
		fmt.Println("rightInt: ", rightInt)
		fmt.Println("operator: ", operator)
	}

	return calc.Calc(leftInt, rightInt, operator)
}

func isNumeric(v string) bool {
	if _, err := strconv.ParseFloat(v, 64); err == nil {
		return true
	}
	return false
}

const (
	PLUS           = "+"
	MINUS          = "-"
	MULTIPLICATION = "*"
	DIVISION       = "/"
)

func (calc *Calculator) Calc(left float64, right float64, operator string) (float64, error) {
	switch operator {
	case PLUS:
		return left + right, nil
	case MINUS:
		return left - right, nil
	case MULTIPLICATION:
		return left * right, nil
	case DIVISION:
		if right == 0 {
			return 0, errors.New("divide by zero:")
		}
		return left / right, nil
	}
	return 0, errors.New("Unknown operator: " + operator)
}
