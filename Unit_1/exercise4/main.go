package main

import "fmt"

type specialNum int

var x specialNum

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42

	fmt.Println(x)
}
