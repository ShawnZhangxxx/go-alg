package main

import "fmt"

const (
	_ = iota + 3
	x
	y
)

func main() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
}