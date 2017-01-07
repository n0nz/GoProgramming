package main

import (
	"fmt"
)

func main() {
	number1:= 200
	number2:= 30
	fmt.Println(sumNumber(number1, number2))
}
func sumNumber(num1, num2 int) int{
	sum := num1 + num2
	return sum
}
