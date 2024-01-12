package main

import "fmt"

func main() {
	const (
		_   = iota // _ == 0
		foo        // foo == 1
		bar        // bar == 2
		far
		fo
		ba
	)

	fmt.Println(foo, bar, far, fo, ba)
}
