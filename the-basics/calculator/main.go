package main

import "fmt"

func main() {
	var operation string
	var num1, num2 int

	fmt.Println("CALCULATOR.COM")
	fmt.Println("==============")
	fmt.Println("Enter the operation you want to perform (+, -, *, /):")

	fmt.Scanf("%s", &operation)

	fmt.Println("Enter the first number:")
	fmt.Scanf("%d", &num1)

	fmt.Println("Enter the second number:")
	fmt.Scanf("%d", &num2)

	switch operation {
	case "+":
		fmt.Println(num1 + num2)

	case "-":
		fmt.Println(num1 - num2)

	case "*":
		fmt.Println(num1 * num2)

	case "/":
		fmt.Println(num1 / num2)
	}
}
