package main

import (
	"fmt"
)

func main() {

	var operator string
	var num1, num2 int

	// newFunction()
	// fmt.Println(data.MaxSpeed)
	
	fmt.Println("CALCULATOR GO 1.0")
	fmt.Println("=================")
	fmt.Println("which operation u want to preform? (add, substract, multiplay, divide)")
	fmt.Scanf("%s\n", &operator)
	fmt.Println("pick first number")
	fmt.Scanf("%d\n", &num1)
	fmt.Println("pick second number")
	fmt.Scanf("%d\n", &num2)
	switch operator {
	case "add":
		fmt.Println( num1 + num2)
	case "substract":
		fmt.Println( num1 - num2)
	case "multiplay":
		fmt.Println( num1 * num2)
	case "divide":
		fmt.Println( num1 / num2)
	}
}