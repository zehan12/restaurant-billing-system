package main

import "fmt"

func main() {
	fmt.Println("Welcome to go restro")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hello, %+v\n", name)
}
