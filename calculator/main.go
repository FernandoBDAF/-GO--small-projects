package main

import "fmt"

func main() {
	var operation string
	var number1, number2 int

	fmt.Println("CALCULATOR GO 1.0")
	fmt.Println(("================="))
	fmt.Println("Enter the operation to perform (add, sub, mult, div): ")
	fmt.Scanf("%s", &operation)
	fmt.Println("Enter the first number: ")
	fmt.Scanf("%d", &number1)
	fmt.Println("Enter the second number: ")
	fmt.Scanf("%d", &number2)
	switch operation {
	case "add":
		fmt.Printf("The result of the sum is: %d\n", number1+number2)
	case "sub":
		fmt.Printf("The result of the subtraction is: %d\n", number1-number2)
	case "mult":
		fmt.Printf("The result of the multiplication is: %d\n", number1*number2)
	case "div":
		fmt.Printf("The result of the division is: %d\n", number1/number2)
	default:
		fmt.Println("Operation not valid")
	}
}