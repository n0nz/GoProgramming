package main

import "fmt"

var g_groceries[]string

func add_grocery(a ...string) {
	fmt.Println("Capacity", cap(g_groceries))
	for _, d := range a {
		g_groceries = append(g_groceries, d)
	}
}

func list_groceries() {
	fmt.Println("Grocery list is as follows:")
	/*
	for i := 0; i < len(g_groceries); i++ {
		fmt.Println(g_groceries[i])
	}
	*/

	for _, d := range g_groceries {
		fmt.Println(d)
	}
}

func main() {
	add_grocery("Hello", "World")
	add_grocery("GoodMorning")
	add_grocery("GoodAfternoon")
	add_grocery("GoodEvening")
	add_grocery("GoodNight")
	add_grocery("NotGoodAnymore", "Easy", "Peasy", "Ofcourse!")
	list_groceries()
}
