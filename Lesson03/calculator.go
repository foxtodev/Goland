package main

import (
	"fmt"
	"strings"
)

func main() {

	var (
		operation  string
		operations string = "+-*/%^"
		a, b, r    float64
	)

	fmt.Print("Enter first operand: ")
	for {
		if _, err := fmt.Scan(&a); err != nil {
			fmt.Print("Incorrect input. Repeat enter: ")
		} else {
			break
		}
	}

	fmt.Print("Enter second operand: ")
	for {
		if _, err := fmt.Scan(&b); err != nil {
			fmt.Print("Incorrect input. Repeat enter: ")
		} else {
			break
		}
	}

	fmt.Print("Enter operation (" + operations + "): ")
	for {
		if _, err := fmt.Scanln(&operation); err != nil {
			fmt.Print("Error. Repeat enter: ")
		} else if len(operation) > 1 || !strings.Contains(operations, operation) {
			fmt.Print("Operation is not supported. Please use " + operations + ": ")
		} else {
			break
		}
	}

	switch operation {
	case "+":
		r = a + b
	case "-":
		r = a - b
	case "*":
		r = a * b
	case "/":
		r = a / b
	case "%":
		r = float64(int(a) % int(b))
	case "^":
		r = 1
		for i := 1; i <= int(b); i++ {
			r = r * a
		}
	default:
		fmt.Println("Err")
	}

	fmt.Println(a, operation, b, "=", r)

}
