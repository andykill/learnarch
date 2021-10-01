package stack

import (
	"fmt"
	"sync"
)

type Stack struct {
	sync.Mutex
	strings []string
	Debug   bool
}

func NewStack() *Stack {
	return &Stack{}
}

func (stack *Stack) Clone() Stack {
	return *stack
}

func (stack *Stack) IsEmpty() bool {
	return len(stack.strings) == 0
}

func (stack *Stack) Pop() string {
	stack.Lock()
	defer stack.Unlock()

	n := len(stack.strings) - 1
	if n < 0 {
		return ""
	}
	str := stack.strings[n]
	stack.strings = stack.strings[:n]

	if stack.Debug {
		fmt.Println("Pop " + str)
	}
	return str
}

func (stack *Stack) Push(str string) {
	stack.Lock()
	defer stack.Unlock()
	stack.strings = append(stack.strings, str)

	if stack.Debug {
		fmt.Println("Push " + str)
	}
}

func (stack *Stack) Strings() []string {
	return stack.strings
}
