package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const INVALID_RESULT = 0

func calculate(op string, a, b int) (int, error) {
	operators := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"*": func(a, b int) int { return a * b },
	}

	if operators[op] == nil {
		err := syntaxError("operator \"%s\" is not supported", op)
		return INVALID_RESULT, err
	}

	return operators[op](a, b), nil
}

func Evaluate(exp string) (int, error) {
	tokens := strings.Split(exp, " ")
	stack := NewStack()

	for _, token := range tokens {
		n, err := strconv.Atoi(token)


		if err == nil {
			stack.Push(n)
			continue
		}

		// tokens is operand
		a, err := stack.Pop()
		if err != nil {
			return INVALID_RESULT, syntaxError("Stack is empty.")
		}

		b, err := stack.Pop()
		if err != nil {
			return INVALID_RESULT, syntaxError("Stack is empty.")
		}

		result, err := calculate(token, a, b)
		if err != nil {
			return INVALID_RESULT, err
		}
		stack.Push(result)
	}

	if stack.Size() != 1 {
		return INVALID_RESULT, syntaxError("")
	}
	result, _ := stack.Pop()

	return result, nil
}

func syntaxError(format string, a ...interface{}) error {
	msg := fmt.Sprintf("syntax error: " + format, a)

	return errors.New(msg)
}
