package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Calculator struct{}

// Methods
func (Calculator) Add(a, b float64) float64 {
	return a + b
}

func (Calculator) Subtract(a, b float64) float64 {
	return a - b
}

func (Calculator) Multiply(a, b float64) float64 {
	return a * b
}

func (Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter calculation (e.g. 2 + 2), or q to quit: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "q" {
			fmt.Println("Goodbye!")
			break
		}

		var a, b float64
		var op string

		n, err := fmt.Sscanf(input, "%f %s %f", &a, &op, &b)
		if err != nil || n != 3 {
			fmt.Println("Invalid input. Please use: number operator number (e.g. 3.5 * 2)")
			continue
		}

		switch op {
		case "+":
			fmt.Println("Result:", a+b)
		case "-":
			fmt.Println("Result:", a-b)
		case "*":
			fmt.Println("Result:", a*b)
		case "/":
			if b == 0 {
				fmt.Println("Error: Cannot divide by zero")
			} else {
				fmt.Println("Result:", a/b)
			}
		default:
			fmt.Println("Unknown operator:", op)
		}
	}
}
