package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Printf("%T\n", x)

	for i, v := range x {
		fmt.Printf("the element at %v is: %v\n", i, v)
	}
}
