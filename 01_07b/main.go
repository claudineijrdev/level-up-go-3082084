package main

import (
	"flag"
	"log"
)

type operatorType int

const (
	openBrackets operatorType = iota
	closedBrackets
	otherOperator
)

var bracketPairs = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

func getOperatorType(op rune) operatorType {
	for ob, cb := range bracketPairs {
		switch op {
		case ob:
			return openBrackets
		case cb:
			return closedBrackets
		}
	}

	return otherOperator
}

type stack struct {
	elems []rune
}

func (s *stack) push(elem rune) {
	s.elems = append(s.elems, elem)
}

func (s *stack) pop() *rune {
	if len(s.elems) == 0 {
		return nil
	}

	n := len(s.elems) - 1
	last := s.elems[n]
	s.elems = s.elems[:n]

	return &last
}

// isBalanced returns whether the given expression
// has balanced brackets.
func isBalanced(expr string) bool {
	s := stack{}
	for _, e := range expr {
		switch getOperatorType(e) {
		case openBrackets:
			s.push(e)
		case closedBrackets:
			last := s.pop()
			if last == nil || bracketPairs[*last] != e {
				return false
			}
		}
	}
	return len(s.elems) == 0
}

// printResult prints whether the expression is balanced.
func printResult(expr string, balanced bool) {
	if balanced {
		log.Printf("%s is balanced.\n", expr)
		return
	}
	log.Printf("%s is not balanced.\n", expr)
}

func main() {
	expr := flag.String("expr", "", "The expression to validate brackets on.")
	flag.Parse()
	printResult(*expr, isBalanced(*expr))
}
