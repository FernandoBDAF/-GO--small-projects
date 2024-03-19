package version

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func CalculatorV2() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("CALCULATOR GO 2.0")
	fmt.Println("=================")
	fmt.Println("Enter the operation to perform (add, sub, mult, div): ")
	scanner.Scan()
	operation := scanner.Text()

	fmt.Println("Enter the first number: ")
	scanner.Scan()
	number1, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Println("Enter the second number: ")
	scanner.Scan()
	number2, _ := strconv.ParseFloat(scanner.Text(), 64)

	switch operation {
	case "add":
		fmt.Printf("The result of the sum is: %f\n", number1+number2)
	case "sub":
		fmt.Printf("The result of the subtraction is: %f\n", number1-number2)
	case "mult":
		fmt.Printf("The result of the multiplication is: %f\n", number1*number2)
	case "div":
		if number2 == 0 {
			fmt.Println("Cannot divide by zero")
		} else {
			fmt.Printf("The result of the division is: %f\n", number1/number2)
		}
	default:
		fmt.Println("Operation not valid")
	}
	println("This is the end of the program. Thanks for using the calculator!/n")
}