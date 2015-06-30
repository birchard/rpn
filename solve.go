package rpn

import (
	"fmt"
	"strconv"
	"strings"
)

// Solve takes an expression as a slice of items
func Solve(expr []string) string {
	if len(expr[0]) == 0 {
		return "empty expression"
	}
	return parse(expr)
}

func parse(expr []string) string {
	pad := new(stack)
	for _, item := range expr {
		operand, err := strconv.ParseFloat(item, 64)
		if err == nil {
			pad.Push(operand)
		} else if !strings.Contains("-+*/", item) {
			return item + " is neither an operator nor an operand."
		} else {
			t1, t2, err := getOperands(pad, item)
			if err != nil {
				return err.Error()
			}

			switch item {
			case "+":
				pad.Push(t2 + t1)
			case "-":
				pad.Push(t2 - t1)
			case "*":
				pad.Push(t2 * t1)
			case "/":
				pad.Push(t2 / t1)
			}
		}
	}
	return strconv.FormatFloat(pad.Pop().(float64), 'f', -1, 64)
}

func getOperands(p *stack, operator string) (float64, float64, error) {
	if p.Len() == 1 {
		switch operator {
		case "-":
			return p.Pop().(float64), 0.0, nil
		default:
			return 0.0, 0.0, fmt.Errorf("unary operator %q is not supported", operator)
		}
	}
	return p.Pop().(float64), p.Pop().(float64), nil
}

// local stack class

type stack struct {
	top  *element
	size int
}

type element struct {
	value interface{} // All types satisfy the empty interface, so we can store anything here.
	next  *element
}

// Return the stack's length
func (s *stack) Len() int {
	return s.size
}

// Push a new element onto the stack
func (s *stack) Push(value interface{}) {
	s.top = &element{value, s.top}
	s.size++
}

// Remove the top element from the stack and return it's value
// If the stack is empty, return nil
func (s *stack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}
