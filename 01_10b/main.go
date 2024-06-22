package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// operators is the map of legal operators and their functions
var operators = map[string]func(x, y float64) float64{
	"+": func(x, y float64) float64 { return x + y },
	"-": func(x, y float64) float64 { return x - y },
	"*": func(x, y float64) float64 { return x * y },
	"/": func(x, y float64) float64 { return x / y },
}

// parseOperand parses a string to a float64
func parseOperand(op string) (float64, error) {
	parsedOp, err := strconv.ParseFloat(op, 64)
	return parsedOp, err
}

// calculate returns the result of a 2 operand mathematical expression
func calculate(expr string) float64 {
	ops := strings.Fields(expr)
	if len(ops) != 3 {
		log.Fatal(fmt.Sprintf("Invalid expression, expected 3 args, found %d", len(ops)))
	}
	left, err := parseOperand(ops[0])
	if err != nil {
		log.Fatal(fmt.Sprintf( "Error while converting value to float64: %s", ops[0]))
	}
	right, err := parseOperand(ops[2])
	if err != nil {
		log.Fatal(fmt.Sprintf( "Error while converting value to float64: %s", ops[2]))
	}

	f, ok := operators[ops[1]]

	if !ok {
		log.Fatal(fmt.Sprintf("Could not select a operation, wrong oparator: %s", ops[1] ))
	}

	result := f(left, right)
	return result
}

func main() {
	expr := flag.String("expr", "",
		"The expression to calculate on, separated by spaces.")
	flag.Parse()
	result := calculate(*expr)
	log.Printf("%s = %.2f\n", *expr, result)
}
