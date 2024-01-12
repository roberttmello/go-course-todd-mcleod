package main

import "fmt"

type ByteSize int

const (
	_   = iota // _ == 0
	foo        // foo == 1
	bar        // bar == 2
	far
	fo
	ba
)

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func main() {

	fmt.Println(foo, bar, far, fo, ba)
	fmt.Println()
	fmt.Printf("%d\t\t\t%b\n", KB, KB)
	fmt.Printf("%d\t\t\t%b\n", MB, MB)
	fmt.Printf("%d\t\t%b\n", GB, GB)
	fmt.Printf("%d\t\t%b\n", TB, TB)
	fmt.Printf("%d\t%b\n", PB, PB)
	fmt.Printf("%d\t%b\n", EB, EB)

}
