package main

import (
	"fmt"
)

func main() {
	var number1 int
	var number2 int
	fmt.Print("Enter number 1: ")
	fmt.Scan(&number1)
	fmt.Print("Enter number 2: ")
	fmt.Scan(&number2)
	fmt.Println(sumNumber(number1, number2))
}
func sumNumber(num1, num2 int) int{
	sum := num1 + num2
	return sum
}
