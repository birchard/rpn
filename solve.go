package rpn

import "strconv"

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
		i, err := strconv.ParseFloat(item, 64)
		if err == nil {
			pad.Push(i)
		} else {
			switch item {
			case "+":
				t1, t2 := additiveOperands(pad)
				pad.Push(t2 + t1)
			case "-":
				t1, t2 := additiveOperands(pad)
				pad.Push(t2 - t1)
			case "*":
				t1, t2 := multiplicativeOperands(pad)
				pad.Push(t2 * t1)
			case "/":
				t1, t2 := multiplicativeOperands(pad)
				pad.Push(t2 / t1)
			default:
				return item + " is neither an operator nor an operand."
			}
		}
	}
	return strconv.FormatFloat(pad.Pop().(float64), 'f', -1, 64)
}

func additiveOperands(p *stack) (float64, float64) {
	if p.Len() > 1 {
		return p.Pop().(float64), p.Pop().(float64)
	}
	return p.Pop().(float64), 0
}

func multiplicativeOperands(p *stack) (float64, float64) {
	if p.Len() > 1 {
		return p.Pop().(float64), p.Pop().(float64)
	}
	return p.Pop().(float64), 0
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
