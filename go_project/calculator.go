package main

import (
	"fmt"
)

var number1 int
var number2 int
var operator string
var Operation int

func main() {
	fmt.Println("Hello user \n wellcome to calculator")
	fmt.Println("+-+-+-+-+-+-+-+-+-+-")
	fmt.Println("\n Enter number1:")
	fmt.Scanln(&number1)
	fmt.Println("Enter number2:")
	fmt.Scanln(&number2)
	fmt.Println("Enter user Operation + - * :")
	fmt.Scanln(&operator)
	switch operator {
	case "+":
		Operation = number1 + number2
		fmt.Println(Operation)
	case "-":
		Operation = number1 - number2
		fmt.Println(Operation)
	case "*":
		Operation = number1 * number2
		fmt.Println(Operation)
	case "/":
		Operation = number1 / number2
		fmt.Println(Operation)
	}

}
