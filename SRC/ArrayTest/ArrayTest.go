package main

import "fmt"

const g_cap int = 5 // total capacity of list
var g_groceries[g_cap]string
var g_len int = 0 // total number of grocery items in current array

func add_grocery(a string) {
	if g_len < g_cap {
		g_groceries[g_len] = a
		g_len++
	} else {
		fmt.Println("Too many items, now we are done for !!")
	}
}

func list_groceries() {
	fmt.Println("Grocery list is as follows:")
	for i := 0; i < g_len; i++ {
		fmt.Println(g_groceries[i])
	}
}

func main() {
	add_grocery("Hello")
	add_grocery("GoodMorning")
	add_grocery("GoodAfternoon")
	add_grocery("GoodEvening")
	add_grocery("GoodNight")
	add_grocery("NotGoodAnymore")
	list_groceries()
}
