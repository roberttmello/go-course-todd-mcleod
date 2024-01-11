package main

import "fmt"

func main() {
	const name, age = "Ted", 28
	fmt.Println("Hello Gophers!")
	fmt.Printf("\n")
	fmt.Printf("1----- %s is %d years old.\n", name, age)
	fmt.Printf("2----- %v is %v years old.\n", name, age)
	fmt.Printf("3----- %s is %d years old. \t The type of the variables is respectively: %T and %T\n", name, age, name, age)
	fmt.Println(`
				Unicode, formally The Unicode Standard, 
				is a text encoding standard maintained by 
				the Unicode Consortium designed to support 
				the use of text written in all of the world's major writing systems. 
				Version 15.1 of the standard defines 
				149813 characters and 161 scripts used 
				in various ordinary, literary, academic, 
				and technical contexts. 
	`)
}
