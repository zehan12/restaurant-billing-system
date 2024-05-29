package main

import (
	"fmt"

	"github.com/zehan12/restaurant-billing-system/menu"
)

func main() {
	fmt.Println("Welcome to go restro")
	var name string
	fmt.Scan(&name)
	// fmt.Printf("Hello, %+v\n", name)
	menu.PrintMenu()

	var items []int
	var input int

	fmt.Println("Enter numbers (enter -1 to stop):")

	for {
		fmt.Scanf("%d", &input)
		if input == -1 {
			break
		}
		items = append(items, input)
	}

	fmt.Println("You entered:", items)

	amount := menu.CalculateAmount(items)
	fmt.Println(amount)

	// generate bill
}
