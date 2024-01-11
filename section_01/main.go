package main

import "fmt"

func main() {
	const name, age = "Ted", 28
	fmt.Println("Hello Gophers!")
	fmt.Printf("\n")
	fmt.Printf("1----- %s is %d years old.\n", name, age)
	fmt.Printf("2----- %v is %v years old.\n", name, age)
	fmt.Printf("3----- %s is %d years old. \t The type of the variables is respectively: %T and %T\n", name, age, name, age)
}
