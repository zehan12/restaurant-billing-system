package menu

import (
	"fmt"
	"math"
)

// MenuItem represents a single menu item with an ID, name, and price.
type MenuItem struct {
	ID    int
	Name  string
	Price float64
}

// Category represents a category of menu items.
type Category struct {
	Name  string
	Items []MenuItem
}

// CreateMenu creates and returns the restaurant menu.
func CreateMenu() []Category {
	return []Category{
		{
			Name: "Appetizers",
			Items: []MenuItem{
				{ID: 1, Name: "Garlic Bread", Price: 5.99},
				{ID: 2, Name: "Bruschetta", Price: 6.99},
				{ID: 3, Name: "Stuffed Mushrooms", Price: 7.99},
			},
		},
		{
			Name: "Salads",
			Items: []MenuItem{
				{ID: 4, Name: "Caesar Salad", Price: 8.99},
				{ID: 5, Name: "Greek Salad", Price: 9.99},
			},
		},
		{
			Name: "Entrees",
			Items: []MenuItem{
				{ID: 6, Name: "Grilled Chicken Alfredo", Price: 14.99},
				{ID: 7, Name: "Beef Lasagna", Price: 13.99},
				{ID: 8, Name: "Vegetarian Pizza", Price: 12.99},
			},
		},
		{
			Name: "Desserts",
			Items: []MenuItem{
				{ID: 9, Name: "Tiramisu", Price: 6.99},
				{ID: 10, Name: "Cheesecake", Price: 5.99},
			},
		},
		{
			Name: "Beverages",
			Items: []MenuItem{
				{ID: 11, Name: "Iced Tea", Price: 2.99},
				{ID: 12, Name: "Lemonade", Price: 3.49},
			},
		},
	}
}

func CalculateAmount(orders []int) string {
	menu := CreateMenu()
	var list []MenuItem
	for i := 0; i < len(menu); i++ {
		for j := 0; j < len(menu[i].Items); j++ {
			list = append(list, menu[i].Items...)
		}
	}

	var amount float64

	for i := 0; i < len(orders); i++ {
		for j := 0; j < len(list); j++ {
			if orders[i] == list[i].ID {
				amount += list[i].Price
			}
		}
	}

	totalAmount := fmt.Sprint("$ ", int(amount))
	return totalAmount
}

func PrintMenu() {
	menu := CreateMenu()
	length := len(menu)
	fmt.Println()
	fmt.Println("Select Item from menu")
	for i := 0; i < length; i++ {
		fmt.Printf("Category: %s\n", menu[i].Name)
		fmt.Println()
		for j := 0; j < len(menu[i].Items); j++ {
			fmt.Printf("#%d %s  $%f\n", menu[i].Items[j].ID, menu[i].Items[j].Name, math.Floor(menu[i].Items[j].Price))
		}
		fmt.Println()
	}
}
