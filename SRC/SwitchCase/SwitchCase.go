package main

import "fmt"

func main() {
	i := 2
	switch i{
	case 1:
		fmt.Println("Case 1")
		fallthrough
	case 2:
		fmt.Println("Case 2")
		fallthrough
	case 3:
		fmt.Println("Case 3")
		fallthrough
	case 4:
		fmt.Println("Case 4")
		fallthrough
	case 5:
		fmt.Println("Case 5")
		fallthrough
	default:
		fmt.Println("Default Case")
	}
}
