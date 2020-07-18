package main

import "fmt"

func main() {
	plus := foo()

	fmt.Println(plus())
	fmt.Println(plus())
	fmt.Println(plus())
	fmt.Println(plus())
	fmt.Println(plus())
	fmt.Println(plus())
	fmt.Println(plus())

	another := foo()

	fmt.Println(another())
	fmt.Println(another())
	fmt.Println(another())
	fmt.Println(another())
	fmt.Println(another())
}

func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
