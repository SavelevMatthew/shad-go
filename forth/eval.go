//go:build !solution

package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type EvaluatorStack = []int

type Operation struct {
	args int
	fn   func(EvaluatorStack) (EvaluatorStack, error)
}

type Evaluator struct {
	stack      EvaluatorStack
	operations map[string]Operation
	aliases    map[string]string
}

func sum(s EvaluatorStack) (EvaluatorStack, error) {
	a := s[len(s)-1]
	b := s[len(s)-2]

	return append(s[:len(s)-2], b+a), nil
}

func sub(s EvaluatorStack) (EvaluatorStack, error) {
	a := s[len(s)-1]
	b := s[len(s)-2]

	return append(s[:len(s)-2], b-a), nil
}

func mul(s EvaluatorStack) (EvaluatorStack, error) {
	a := s[len(s)-1]
	b := s[len(s)-2]

	return append(s[:len(s)-2], b*a), nil
}

func div(s EvaluatorStack) (EvaluatorStack, error) {
	a := s[len(s)-1]
	b := s[len(s)-2]

	if a == 0 {
		return s, errors.New("division by zero")
	}

	return append(s[:len(s)-2], b/a), nil
}

func dup(s EvaluatorStack) (EvaluatorStack, error) {
	a := s[len(s)-1]

	return append(s, a), nil
}

func drop(s EvaluatorStack) (EvaluatorStack, error) {
	return s[:len(s)-1], nil
}

func swap(s EvaluatorStack) (EvaluatorStack, error) {
	a := s[len(s)-1]
	b := s[len(s)-2]

	return append(s[:len(s)-2], a, b), nil
}

func over(s EvaluatorStack) (EvaluatorStack, error) {
	b := s[len(s)-2]

	return append(s, b), nil
}

// NewEvaluator creates evaluator.
func NewEvaluator() *Evaluator {
	return &Evaluator{
		stack: make([]int, 0),
		operations: map[string]Operation{
			"+":    {2, sum},
			"-":    {2, sub},
			"/":    {2, div},
			"*":    {2, mul},
			"dup":  {1, dup},
			"drop": {1, drop},
			"swap": {2, swap},
			"over": {2, over},
		},
		aliases: make(map[string]string),
	}
}

func (e *Evaluator) checkMinArgs(min int) error {
	if len(e.stack) < min {
		return fmt.Errorf("at least %v args required, but %v was found", min, len(e.stack))
	}

	return nil
}

func (e *Evaluator) proccessCommand(row string, isPrimitive bool) ([]int, error) {
	parts := strings.Split(row, " ")

	if parts[0] == ":" && parts[len(parts)-1] == ";" && len(parts) > 2 {
		name := strings.ToLower(parts[1])
		_, err := strconv.Atoi(name)
		if err == nil {
			return e.stack, fmt.Errorf("%v cannot be command name", name)
		}

		alias := make([]string, 0, len(parts)-3)
		for _, part := range parts[2 : len(parts)-1] {
			_, err := strconv.Atoi(part)
			if err == nil {
				alias = append(alias, part)
				continue
			}

			low := strings.ToLower(part)

			if a, ok := e.aliases[low]; ok {
				alias = append(alias, a)
				continue
			}

			if _, ok := e.operations[low]; ok {
				alias = append(alias, part)
				continue
			}

			return e.stack, fmt.Errorf("unknown operation: %v", low)
		}

		e.aliases[name] = strings.Join(alias, " ")
		return e.stack, nil
	}

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err == nil {
			e.stack = append(e.stack, num)
			continue
		}

		low := strings.ToLower(part)

		if alias, ok := e.aliases[low]; ok && !isPrimitive {
			e.stack, err = e.proccessCommand(alias, true)
			if err != nil {
				return e.stack, err
			}

			continue
		}

		if op, ok := e.operations[low]; ok {
			err := e.checkMinArgs(op.args)
			if err != nil {
				return e.stack, err
			}
			e.stack, err = op.fn(e.stack)
			if err != nil {
				return e.stack, err
			}

			continue
		}

		return e.stack, fmt.Errorf("unknown operation: %v", low)
	}

	return e.stack, nil
}

// Process evaluates sequence of words or definition.
//
// Returns resulting stack state and an error.
func (e *Evaluator) Process(row string) ([]int, error) {
	return e.proccessCommand(row, false)
}
