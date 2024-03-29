package main

import (
	"fmt"
)

// EVERYTHING declared or imported in your code most be used!!!

// isNotGoingWork := "Teste"
var thatsFine string = "Good"
var test = 10

func main() {

	fmt.Println(thatsFine)
	fmt.Println(test)

	// Implicit type declarations
	var age, year, _, name, date, isValid, price = 28, 2024, "value ignored", "Robert", "11/01/2024", true, 6547.698
	fmt.Println(age, year, name, date, isValid, price)

	// Shorthand declaration
	magicNumber := 42
	fmt.Println(magicNumber)

	// Explicit type declarations
	var age1 int = 68
	var year1 int = 2024
	var name1 string = "Rob Pike"
	var date1 string = "11/01/2024"
	var isValid1 bool = false
	var price1 float64 = 25.69

	fmt.Println(age1, year1, name1, date1, isValid1, price1)

	// Zero values will be assigned
	var age2 int
	var year2 int
	var name2 string
	var date2 string
	var isValid2 bool
	var price2 float64

	fmt.Println(age2, year2, name2, date2, isValid2, price2)

	// Print these values as both binary and hexadecimal
	a, b, c, d, e, f := 10, 11, 12, 13, 14, 15
	fmt.Printf("Binary format: %b %b %b %b %b %b\n", a,b,c,d,e,f)
	fmt.Printf("Hexadecimal format: %x %x %x %x %x %x\n", a,b,c,d,e,f)
	
	// Conversions
	z := 42.0
	var m float32 = 48.75
	z = float64(m)
	fmt.Println(z)

	// Iota

	const (
		_ = iota
		foo
		bar
		far 
		fo 
		ba
	)

	fmt.Println(foo, bar, far, fo, ba)


}
